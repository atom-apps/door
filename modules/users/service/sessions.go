package service

import (
	"context"
	"time"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-providers/uuid"

	"github.com/jinzhu/copier"
)

// @provider
type SessionService struct {
	uuid       *uuid.Generator
	hash       *md5.Hash
	sessionDao *dao.SessionDao
	tokenDao   *dao.TokenDao
	userDao    *dao.UserDao
}

func (svc *SessionService) DecorateItem(model *models.Session, id int) *dto.SessionItem {
	dtoItem := &dto.SessionItem{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		UserID:    model.UserID,
		SessionID: model.SessionID,
		ExpireAt:  model.ExpireAt,
		User:      nil,
		Tokens:    nil,
	}

	user, err := svc.userDao.GetByID(context.Background(), model.UserID)
	if err == nil {
		dtoItem.User = user
	}

	tokens, err := svc.tokenDao.GetBySessionIDWithoutScope(context.Background(), model.ID)
	if err == nil {
		dtoItem.Tokens = tokens
	}

	return dtoItem
}

func (svc *SessionService) GetByID(ctx context.Context, id int64) (*models.Session, error) {
	return svc.sessionDao.GetByID(ctx, id)
}

func (svc *SessionService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.SessionListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Session, error) {
	return svc.sessionDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *SessionService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.SessionListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Session, int64, error) {
	return svc.sessionDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *SessionService) CreateFromModel(ctx context.Context, model *models.Session) error {
	return svc.sessionDao.Create(ctx, model)
}

// Create
func (svc *SessionService) Create(ctx context.Context, body *dto.SessionForm) error {
	model := &models.Session{}
	_ = copier.Copy(model, body)
	return svc.sessionDao.Create(ctx, model)
}

// CreateForUser
func (svc *SessionService) CreateForUser(ctx context.Context, userID int64, sessionID string) (*models.Session, error) {
	sess, _ := svc.sessionDao.GetBySessionID(ctx, sessionID)
	if sess != nil {
		return sess, nil
	}

	model := &models.Session{
		UserID:    userID,
		SessionID: sessionID,
		ExpireAt:  time.Now().Add(time.Duration(common.SessionExpireDuration*24) * time.Hour),
	}
	if err := svc.CreateFromModel(ctx, model); err != nil {
		return nil, err
	}
	return model, nil
}

func (svc *SessionService) GetBySessionID(ctx context.Context, sessionID string) (*models.Session, error) {
	return svc.sessionDao.GetBySessionID(ctx, sessionID)
}

// GenerateSessionID
func (svc *SessionService) GenerateSessionID() string {
	return svc.hash.Hash(svc.uuid.MustGenerate())
}

// DeleteBySessionID
func (svc *SessionService) DeleteBySessionID(ctx context.Context, sessID int64) error {
	return svc.sessionDao.Transaction(func() error {
		if err := svc.sessionDao.DeletePermanently(ctx, sessID); err != nil {
			return err
		}

		if err := svc.tokenDao.DeleteBySessionID(ctx, sessID); err != nil {
			return err
		}
		return nil
	})
}
