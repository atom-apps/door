package dao

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	
	if err := container.Container.Provide(NewUserDao); err!=nil {
		return err
	}
	
	if err := container.Container.Provide(NewRoleDao); err!=nil {
		return err
	}
	return nil
}
