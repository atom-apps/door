// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/modules/service/controller"
	"github.com/atom-apps/door/modules/service/dto"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeSendController(engine fiber.Router, controller *controller.SendController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Post(strings.TrimPrefix("/v1/services/send/sms", groupPrefix), Func1(controller.Sms, Body[dto.SendVerifyCodeForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/services/send/email", groupPrefix), Func1(controller.Email, Body[dto.SendVerifyCodeForm](BodyParamError)))
}
