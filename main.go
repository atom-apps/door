//go:generate atomctl gen enum
//go:generate atomctl gen routes
//go:generate atomctl gen provider
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"log"

	moduleAuth "github.com/atom-apps/door/modules/auth"
	moduleService "github.com/atom-apps/door/modules/service"
	moduleUser "github.com/atom-apps/door/modules/user"

	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-apps/door/providers/jwt"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-apps/door/providers/oauth"
	databasePostgres "github.com/atom-providers/database-postgres"
	redis "github.com/atom-providers/database-redis"
	service "github.com/atom-providers/service-httpgrpc"
	"github.com/atom-providers/swagger"
	"github.com/atom-providers/uuid"
	_ "github.com/go-gormigrate/gormigrate/v2"
	"github.com/rogeecn/atom"
	"github.com/spf13/cobra"
)

func main() {
	providers := service.
		Default(
			redis.DefaultProvider(),
			md5.DefaultProvider(),
			bcrypt.DefaultProvider(),
			uuid.DefaultProvider(),
			oauth.DefaultProvider(),
			swagger.DefaultProvider(),
			query.DefaultProvider(),
			jwt.DefaultProvider(),
			databasePostgres.DefaultProvider(),
		).
		With(
			moduleUser.Providers(),
			moduleAuth.Providers(),
			moduleService.Providers(),
		)

	opts := []atom.Option{
		atom.Name("door"),
		atom.RunE(func(cmd *cobra.Command, args []string) error {
			return service.Serve()
		}),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}
