//go:generate atomctl gen routes
//go:generate atomctl gen enum
//go:generate atomctl gen provider
//go:generate swag fmt
//go:generate gofumpt -l -w .
//go:generate swag init -ot json --parseDependency --generatedTime
package main

import (
	"log"

	"github.com/atom-apps/door/modules/boot"

	moduleAuth "github.com/atom-apps/door/modules/auth"
	moduleService "github.com/atom-apps/door/modules/service"
	moduleSystem "github.com/atom-apps/door/modules/systems"
	moduleUser "github.com/atom-apps/door/modules/users"

	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/captcha"
	"github.com/atom-providers/casbin"
	database "github.com/atom-providers/database-mysql"
	redis "github.com/atom-providers/database-redis"
	"github.com/atom-providers/hashids"
	"github.com/atom-providers/jwt"
	service "github.com/atom-providers/service-http"
	"github.com/atom-providers/swagger"
	"github.com/atom-providers/uuid"
	_ "github.com/go-gormigrate/gormigrate/v2"
	"github.com/rogeecn/atom"
	"github.com/spf13/cobra"
)

func main() {
	providers := service.
		Default(
			hashids.DefaultProvider(),
			redis.DefaultProvider(),
			casbin.DefaultProvider(),
			md5.DefaultProvider(),
			captcha.DefaultProvider(),
			bcrypt.DefaultProvider(),
			uuid.DefaultProvider(),
			oauth.DefaultProvider(),
			swagger.DefaultProvider(),
			query.DefaultProvider(),
			jwt.DefaultProvider(),
			database.DefaultProvider(),
		).
		With(boot.Providers()).
		With(
			moduleUser.Providers(),
			moduleAuth.Providers(),
			moduleService.Providers(),
			moduleSystem.Providers(),
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
