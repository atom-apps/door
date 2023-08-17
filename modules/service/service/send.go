package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/atom-providers/app"
	"github.com/atom-providers/log"
	redis "github.com/redis/go-redis/v9"
)

// @provider
type SendService struct {
	cache redis.Cmdable
	app   *app.Config
}

func (svc *SendService) GenerateRandomCode(ctx context.Context) string {
	if svc.app.Mode == app.AppModeDevelopment {
		return "123123"
	}

	// 生成一个 6 位数
	randomNum := rand.Intn(900000) + 100000

	// 如果开头是0，重新生成随机数，直到开头不是0
	for randomNum/100000 == 0 {
		randomNum = rand.Intn(900000) + 100000
	}

	return fmt.Sprintf("%d", randomNum)
}

func (svc *SendService) SendEmailCode(ctx context.Context, target string) error {
	code := svc.GenerateRandomCode(ctx)
	svc.cache.Set(ctx, consts.CacheKeyRegisterCode.With(target), code, time.Minute*10)

	log.Debugf("send email verify code(%s) to %s", code, target)
	return nil
}

func (svc *SendService) SendSmsCode(ctx context.Context, target string) error {
	code := svc.GenerateRandomCode(ctx)
	svc.cache.Set(ctx, consts.CacheKeyRegisterCode.With(target), code, time.Minute*10)

	log.Debugf("send sms verify code(%s) to %s", code, target)
	return nil
}

func (svc *SendService) VerifyCode(ctx context.Context, target, code string) bool {
	cacheCode, err := svc.cache.Get(ctx, consts.CacheKeyRegisterCode.With(target)).Result()
	if err != nil {
		log.Error(err)
		return false
	}
	return cacheCode == code
}
