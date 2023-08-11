package service

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func() (*SendService, error) {
		obj := &SendService{}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
