package service

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/errorx"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-providers/hashids"
	"github.com/atom-providers/log"
	"github.com/samber/lo"

	"github.com/jinzhu/copier"
)

// @provider
type UserService struct {
	userDao           *dao.UserDao
	hashID            *hashids.HashID
	hash              *bcrypt.Hash
	permissionSvc     *PermissionService
	userTenantRoleSvc *UserTenantRoleService
}

func (svc *UserService) DecorateItem(model *models.User, id int) *dto.UserItem {
	tenants, err := svc.userTenantRoleSvc.GetTenantsByUserID(context.Background(), model.ID)
	if err != nil {
		log.Warnf("get tenants of user %d failed: %v", model.ID, err)
	}

	dtoItem := &dto.UserItem{
		ID:            model.ID,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		UUID:          model.UUID,
		Username:      model.Username,
		Email:         model.Email,
		EmailVerified: model.EmailVerified,
		Phone:         model.Phone,
		DisplayName:   model.DisplayName,
		Avatar:        model.Avatar,
		Status:        model.Status,
		TenantRoles: lo.FilterMap(tenants, func(tenant *models.Tenant, _ int) (*dto.UserItemTenantRole, bool) {
			role, err := svc.userTenantRoleSvc.GetRoleOfTenantUser(context.Background(), tenant.ID, model.ID)
			if err != nil {
				log.Warnf("get role of tenant %d user %d failed: %v", tenant.ID, model.ID, err)
				return nil, false
			}

			return &dto.UserItemTenantRole{Role: role, Tenant: tenant}, true
		}),
	}

	return dtoItem
}

func (svc *UserService) GetByID(ctx context.Context, id uint64) (*models.User, error) {
	return svc.userDao.GetByID(ctx, id)
}

func (svc *UserService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, error) {
	return svc.userDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *UserService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, int64, error) {
	return svc.userDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *UserService) CreateFromModel(ctx context.Context, model *models.User) error {
	return svc.userDao.Create(ctx, model)
}

// Create
func (svc *UserService) Create(ctx context.Context, body *dto.UserForm) error {
	model := &models.User{}
	_ = copier.Copy(model, body)
	return svc.userDao.Create(ctx, model)
}

// Update
func (svc *UserService) Update(ctx context.Context, id uint64, body *dto.UserForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.userDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *UserService) UpdateFromModel(ctx context.Context, model *models.User) error {
	return svc.userDao.Update(ctx, model)
}

// Delete
func (svc *UserService) Delete(ctx context.Context, id uint64) error {
	return svc.userDao.Transaction(func() error {
		if err := svc.userDao.Delete(ctx, id); err != nil {
			return err
		}

		return svc.userTenantRoleSvc.DeleteByUserID(ctx, id)
	})
}

// FindByEmail
func (svc *UserService) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return svc.userDao.FindByEmail(ctx, email)
}

// FindByPhone
func (svc *UserService) FindByPhone(ctx context.Context, phone string) (*models.User, error) {
	return svc.userDao.FindByPhone(ctx, phone)
}

// FindByUsername
func (svc *UserService) FindByUsername(ctx context.Context, userName string) (*models.User, error) {
	return svc.userDao.FindByUsername(ctx, userName)
}

// IsEmailValid
func (svc *UserService) IsEmailValid(ctx context.Context, email string) bool {
	regex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

// IsPhoneValid
func (svc *UserService) IsPhoneValid(ctx context.Context, phone string) bool {
	regex := `^(?:(?:\+|00)86)?1[3-9]\d{9}$`
	match, _ := regexp.MatchString(regex, phone)
	return match
}

// IsUsernameValid
func (svc *UserService) IsUsernameValid(ctx context.Context, username string) bool {
	regex := `^[a-z0-9_]{3,20}$`
	match, _ := regexp.MatchString(regex, username)
	return match
}

// GetByUsernameOrEmailOrPhone
func (svc *UserService) GetByUsernameOrEmailOrPhone(ctx context.Context, input string) (*models.User, error) {
	if input == "" {
		return nil, errorx.ErrorUsernameOrPasswordInvalid
	}

	return svc.userDao.GetByUsernameOrEmailOrPhone(ctx, input)
}

// GetByPhone
func (svc *UserService) GetByPhone(ctx context.Context, input string) (*models.User, error) {
	if input == "" {
		return nil, errorx.ErrorUsernameOrPasswordInvalid
	}

	return svc.userDao.GetByPhone(ctx, input)
}

// GetByEmail
func (svc *UserService) GetByEmail(ctx context.Context, input string) (*models.User, error) {
	if input == "" {
		return nil, errorx.ErrorUsernameOrPasswordInvalid
	}

	return svc.userDao.GetByEmail(ctx, input)
}

// GetUserIDHashToken
func (svc *UserService) GetUserIDHashToken(ctx context.Context, user *models.User) (string, error) {
	salt := common.RandomString(10)
	hashid, err := svc.hashID.EncodeWithSalt(salt, int64(user.ID))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", hashid, salt), nil
}

// GetUserFromHashToken
func (svc *UserService) GetUserFromHashToken(ctx context.Context, token string) (*models.User, error) {
	sections := strings.Split(token, "/")
	hashid, salt := sections[0], sections[1]

	ids, err := svc.hashID.DecodeWithSalt(salt, hashid)
	if err != nil {
		return nil, err
	}

	return svc.userDao.GetByID(ctx, uint64(ids[0]))
}

func (svc *UserService) ResetPassword(ctx context.Context, user *models.User, password string) error {
	user.Password = svc.hash.Hash(password)
	return svc.UpdateFromModel(ctx, user)
}
