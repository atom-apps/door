package doorClient

import (
	"github.com/imroc/req/v3"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

type Client struct {
	client *req.Client

	Permission *permission
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func() (*Client, error) {
		client := &Client{
			client: req.C(),
		}

		client.Permission = &permission{client.client}

		return client, nil
	}, o.DiOptions()...)
}

func (c *Client) SetBaseURL(url string) *Client {
	c.client = c.client.SetBaseURL(url)
	return c
}
