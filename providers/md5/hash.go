package md5

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Md5"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Hash struct{}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Hash
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func() *Hash {
		return &config
	}, o.DiOptions()...)
}

func (c *Hash) Hash(raw string) string {
	hash := md5.New()
	hash.Write([]byte(raw))
	return hex.EncodeToString(hash.Sum(nil))
}

func (c *Hash) Compare(plain, hashed string) bool {
	return c.Hash(plain) == hashed
}
