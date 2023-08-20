package seeders

import (
	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type PermissionRulesSeeder struct{}

func NewPermissionRulesSeeder() contracts.Seeder {
	return &PermissionRulesSeeder{}
}

func (s *PermissionRulesSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.PermissionRule{}).TableName(nil))
	db.CreateInBatches(s.Generate(faker, 1), 10)
}

func (s *PermissionRulesSeeder) Generate(faker *gofakeit.Faker, idx int) []models.PermissionRule {
	return []models.PermissionRule{
		s.genPolicyModel("2", "1", "/v1/users/tenants", "GET"),
		s.genPolicyModel("2", "2", "/v1/users/tenants", "GET"),

		s.genRoleModel("2", "2", "1"),
		s.genRoleModel("3", "2", "2"),
	}
}

// roleID, tenantID, path, action}
func (svc *PermissionRulesSeeder) genPolicyModel(args ...string) models.PermissionRule {
	return models.PermissionRule{Ptype: "p", V0: args[0], V1: args[1], V2: args[2], V3: args[3], V4: "", V5: ""}
}

// userID, roleID, tenantID
func (svc *PermissionRulesSeeder) genRoleModel(args ...string) models.PermissionRule {
	return models.PermissionRule{Ptype: "g", V0: args[0], V1: args[1], V2: args[2], V3: "", V4: "", V5: ""}
}
