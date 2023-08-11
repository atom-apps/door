package bcrypt

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"golang.org/x/crypto/bcrypt"
)

const DefaultPrefix = "Bcrypt"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Hash struct {
	Cost int
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Hash
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	if config.Cost == 0 {
		config.Cost = bcrypt.DefaultCost
	}

	return container.Container.Provide(func() *Hash {
		return &config
	}, o.DiOptions()...)
}

func (c *Hash) Hash(raw string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(raw), c.Cost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (c *Hash) Compare(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
