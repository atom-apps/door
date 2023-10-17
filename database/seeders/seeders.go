package seeders

import "github.com/rogeecn/atom/contracts"

var Seeders = []contracts.SeederProvider{
	NewUserSeeder,
	NewTenantSeeder,
	NewRoleSeeder,
	NewPermissionsSeeder,
	NewUserTenantRoleSeeder,
	NewRouteSeeder,
	NewDictionarySeeder,
	NewMenuSeeder,
	NewLocationSeeder,
	NewUserAddressSeeder,

	// seeders
	NewDriverSeeder,
	NewFilesystemSeeder,

	// posts
	NewTagSeeder,
	NewBookSeeder,
	NewChapterSeeder,
	NewArticleSeeder,
}
