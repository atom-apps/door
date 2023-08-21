package doorClient

import "github.com/imroc/req/v3"

type permission struct {
	client *req.Client
}

func (c *permission) Check(path, method string) bool {
	body := struct {
		Path   string `json:"path"`
		Method string `json:"method"`
	}{
		Path:   path,
		Method: method,
	}

	resp, err := c.client.R().SetBody(body).Post("/v1/permission/check")
	if err != nil {
		return false
	}

	return resp.IsSuccessState()
}
