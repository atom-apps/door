package service

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	
	if err := container.Container.Provide(NewUserService); err!=nil {
		return err
	}
	
	if err := container.Container.Provide(NewRoleService); err!=nil {
		return err
	}
	return nil
}
