package service

import (
	"context"
	"time"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/user/dao"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/providers/jwt"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/uuid"
	"github.com/samber/lo"
	"golang.org/x/oauth2"

	"github.com/jinzhu/copier"
)

// @provider
type TokenService struct {
	hash     *md5.Hash
	uuid     *uuid.Generator
	jwt      *jwt.JWT
	tokenDao *dao.TokenDao
	userDao  *dao.UserDao
}

func (svc *TokenService) DecorateItem(model *models.Token, id int) *dto.TokenItem {
	var dtoItem *dto.TokenItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *TokenService) GetByID(ctx context.Context, id int64) (*models.Token, error) {
	return svc.tokenDao.GetByID(ctx, id)
}

func (svc *TokenService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TokenListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Token, error) {
	return svc.tokenDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *TokenService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TokenListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
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
func (svc *TokenService) Update(ctx context.Context, id int64, body *dto.TokenForm) error {
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
func (svc *TokenService) Delete(ctx context.Context, id int64) error {
	return svc.tokenDao.Delete(ctx, id)
}

// getClaims
func (svc *TokenService) getClaims(ctx context.Context, userID int64) *jwt.Claims {
	return svc.jwt.CreateClaims(jwt.BaseClaims{
		UserID: userID,
	})
}

// CreateForUser
func (svc *TokenService) CreateForUser(ctx context.Context, userID, sessID int64, app *oauth.App) (*models.Token, error) {
	m, _ := svc.tokenDao.GetBySessionID(ctx, sessID, app.Name)
	if m != nil {
		return m, nil
	}

	claim := svc.getClaims(ctx, userID)
	token, err := svc.jwt.WithExpireTime(app.TokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	refreshToken, err := svc.jwt.WithExpireTime(app.RefreshTokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	model := &models.Token{
		UserID:        userID,
		SessionID:     sessID,
		AccessToken:   token,
		RefreshToken:  refreshToken,
		ExpireAt:      time.Now().Add(lo.Must1(time.ParseDuration(app.TokenDuration))),
		Scope:         app.Name,
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
func (svc *TokenService) GetByToken(ctx context.Context, token, appName string) (*models.Token, error) {
	return svc.tokenDao.GetByToken(ctx, token, appName)
}

// GetByUserID
func (svc *TokenService) GetByUserID(ctx context.Context, userID int64, appName string) (*models.Token, error) {
	return svc.tokenDao.GetByUserID(ctx, userID, appName)
}

// GetByRefreshToken
func (svc *TokenService) GetByRefreshToken(ctx context.Context, refreshToken, appName string) (*models.Token, error) {
	return svc.tokenDao.GetByRefreshToken(ctx, refreshToken, appName)
}

// RefreshToken
func (svc *TokenService) RefreshToken(ctx context.Context, token *models.Token, app *oauth.App) (*models.Token, error) {
	claim := svc.getClaims(ctx, token.UserID)
	accessToken, err := svc.jwt.WithExpireTime(app.TokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	refreshToken, err := svc.jwt.WithExpireTime(app.RefreshTokenDuration).CreateToken(claim)
	if err != nil {
		return nil, err
	}

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	token.ExpireAt = time.Now().Add(lo.Must1(time.ParseDuration(app.TokenDuration)))

	if err := svc.UpdateFromModel(ctx, token); err != nil {
		return nil, err
	}
	return token, nil
}

// GetByCode
func (svc *TokenService) GetOAuthTokenByCode(ctx context.Context, code, scope string) (*oauth2.Token, error) {
	model, err := svc.tokenDao.GetByCode(ctx, code, scope)
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
func (svc *TokenService) GetBySessionID(ctx context.Context, sessionID int64, scope string) (*models.Token, error) {
	return svc.tokenDao.GetBySessionID(ctx, sessionID, scope)
}
