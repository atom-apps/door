package service

import (
	"context"
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/uuid"
	"github.com/samber/lo"
	"golang.org/x/oauth2"

	"github.com/jinzhu/copier"
)

// @provider
type TokenService struct {
	hash              *md5.Hash
	uuid              *uuid.Generator
	jwt               *jwt.JWT
	auth              *oauth.Auth
	tokenDao          *dao.TokenDao
	userDao           *dao.UserDao
	sessionDao        *dao.SessionDao
	roleDao           *dao.RoleDao
	UserTenantRoleSvc *UserTenantRoleService
}

func (svc *TokenService) DecorateItem(model *models.Token, id int) *dto.TokenItem {
	item := &dto.TokenItem{
		CreatedAt:     model.CreatedAt,
		UserID:        model.UserID,
		AccessToken:   model.AccessToken,
		RefreshToken:  model.RefreshToken,
		Scope:         model.Scope,
		TokenType:     model.TokenType,
		CodeChallenge: model.CodeChallenge,
		Code:          model.Code,
		CodeExpireAt:  model.CodeExpireAt,
		SessionID:     model.SessionID,
		ExpireAt:      model.ExpireAt,
		Used:          model.Used,
		Session:       nil,
	}

	sess, err := svc.sessionDao.GetByID(context.Background(), model.SessionID)
	if err == nil {
		item.Session = sess
	}

	return item
}

func (svc *TokenService) GetByID(ctx context.Context, id uint64) (*models.Token, error) {
	return svc.tokenDao.GetByID(ctx, id)
}

func (svc *TokenService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TokenListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Token, error) {
	return svc.tokenDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *TokenService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TokenListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Token, int64, error) {
	return svc.tokenDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *TokenService) CreateFromModel(ctx context.Context, model *models.Token) error {
	return svc.tokenDao.Create(ctx, model)
}

// Create
func (svc *TokenService) Create(ctx context.Context, body *dto.TokenForm) error {
	model := &models.Token{}
	_ = copier.Copy(model, body)
	return svc.tokenDao.Create(ctx, model)
}

// Update
func (svc *TokenService) Update(ctx context.Context, id uint64, body *dto.TokenForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.tokenDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *TokenService) UpdateFromModel(ctx context.Context, model *models.Token) error {
	return svc.tokenDao.Update(ctx, model)
}

// Delete
func (svc *TokenService) Delete(ctx context.Context, id uint64) error {
	return svc.tokenDao.Delete(ctx, id)
}

// getClaims
func (svc *TokenService) getClaims(ctx context.Context, userID, tenantID, roleID uint64, role string) *jwt.Claims {
	return svc.jwt.CreateClaims(jwt.BaseClaims{
		UserID:   userID,
		TenantID: tenantID,
		RoleID:   roleID,
		Role:     role,
	})
}

// CreateForUser
func (svc *TokenService) CreateForUser(ctx context.Context, userID, tenantID, sessID uint64) (*models.Token, error) {
	m, _ := svc.tokenDao.GetBySessionID(ctx, sessID)
	if m != nil {
		return m, nil
	}

	var err error
	var role *models.Role
	if tenantID != 0 {
		role, err = svc.UserTenantRoleSvc.GetRoleOfTenantUser(ctx, tenantID, userID)
		if err != nil {
			return nil, err
		}
	} else {
		role, err = svc.roleDao.GetBySlug(ctx, jwt.RoleSuperAdmin.String())
		if err != nil {
			return nil, err
		}
	}

	claim := svc.getClaims(ctx, userID, tenantID, role.ID, role.Slug)
	token, err := svc.jwt.WithExpireTime(svc.auth.TokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	refreshToken, err := svc.jwt.WithExpireTime(svc.auth.RefreshTokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	model := &models.Token{
		TenantID:      tenantID,
		UserID:        userID,
		SessionID:     sessID,
		AccessToken:   token,
		RefreshToken:  refreshToken,
		ExpireAt:      time.Now().Add(lo.Must1(time.ParseDuration(svc.auth.TokenDuration))),
		Scope:         "",
		TokenType:     consts.TokenTypeBearer,
		CodeChallenge: "",
		Code:          svc.hash.Hash(svc.uuid.MustGenerate()),
		CodeExpireAt:  time.Now().Add(time.Minute),
		Used:          false,
	}

	if err := svc.CreateFromModel(ctx, model); err != nil {
		return nil, err
	}
	return model, nil
}

// GetByToken
func (svc *TokenService) GetByToken(ctx context.Context, token string) (*models.Token, error) {
	return svc.tokenDao.GetByToken(ctx, token)
}

// GetByUserID
func (svc *TokenService) GetByUserID(ctx context.Context, userID uint64) (*models.Token, error) {
	return svc.tokenDao.GetByUserID(ctx, userID)
}

// GetByRefreshToken
func (svc *TokenService) GetByRefreshToken(ctx context.Context, refreshToken string) (*models.Token, error) {
	return svc.tokenDao.GetByRefreshToken(ctx, refreshToken)
}

// RefreshToken
func (svc *TokenService) RefreshToken(ctx context.Context, token *models.Token) (*models.Token, error) {
	role, err := svc.UserTenantRoleSvc.GetRoleOfTenantUser(ctx, token.TenantID, token.UserID)
	if err != nil {
		return nil, err
	}

	claim := svc.getClaims(ctx, token.UserID, token.TenantID, role.ID, role.Slug)
	accessToken, err := svc.jwt.WithExpireTime(svc.auth.TokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	refreshToken, err := svc.jwt.WithExpireTime(svc.auth.RefreshTokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	token.ExpireAt = time.Now().Add(lo.Must1(time.ParseDuration(svc.auth.TokenDuration)))

	if err := svc.UpdateFromModel(ctx, token); err != nil {
		return nil, err
	}
	return token, nil
}

// GetByCode
func (svc *TokenService) GetOAuthTokenByCode(ctx context.Context, code string) (*oauth2.Token, error) {
	model, err := svc.tokenDao.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	model.Used = true
	if err := svc.UpdateFromModel(ctx, model); err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  model.AccessToken,
		RefreshToken: model.RefreshToken,
		Expiry:       model.ExpireAt,
	}, nil
}

// GetBySessionID
func (svc *TokenService) GetBySessionID(ctx context.Context, sessionID uint64) (*models.Token, error) {
	return svc.tokenDao.GetBySessionID(ctx, sessionID)
}
