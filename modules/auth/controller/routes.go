package controller

import (
	"encoding/json"

	"github.com/atom-apps/door/docs"
	"github.com/atom-apps/door/modules/auth/dto"
	"github.com/gofiber/fiber/v2"
)

// @provider
type RoutesController struct{}

// List
//
//	@Summary		Signup
//	@Description	Signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.ExchangeTokenByCodeForm
//	@Router			/auth/routes [get]
func (c *RoutesController) List(ctx *fiber.Ctx) ([]*dto.Route, error) {
	var doc *dto.SwaggerDoc
	err := json.Unmarshal([]byte(docs.SwaggerSpec), &doc)
	if err != nil {
		return nil, err
	}
	return doc.ToRoues(), nil
}
