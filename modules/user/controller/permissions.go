package controller

import (
	"github.com/atom-apps/door/modules/user/service"
)

// @provider
type PermissionRuleController struct {
	permissionRuleSvc *service.PermissionRuleService
}
