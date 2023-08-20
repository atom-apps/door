package seeders

import "github.com/rogeecn/atom/contracts"

var Seeders = []contracts.SeederProvider{
	NewUserSeeder,
	NewTenantSeeder,
	NewTenantUserSeeder,
	NewRoleSeeder,
	NewRoleUserSeeder,
	NewPermissionRulesSeeder,
}
