package main

import (
	"authsvc/internal/application/auth/login"
	"authsvc/internal/application/auth/register"
	"authsvc/internal/infrastructure/crosscutting/bcrypt"
	"authsvc/internal/infrastructure/crosscutting/pgsql_client"
	viperconfig "authsvc/internal/infrastructure/crosscutting/viper_config"
	"authsvc/internal/infrastructure/paseto_tokens"
	sessionpgsql "authsvc/internal/infrastructure/repository/session/pgsql"
	userpgsql "authsvc/internal/infrastructure/repository/user/pgsql"
	"authsvc/internal/infrastructure/tx/pgsqltx"
	httprest "authsvc/internal/interface/http_rest"
	"authsvc/internal/interface/http_rest/auth/post_login"
	"authsvc/internal/interface/http_rest/auth/post_register"
	"time"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		// config initialization
		fx.Provide(viperconfig.New),
		fx.Provide(bcrypt.New),
		// postgres connection pool
		fx.Provide(pgsql_client.New),
		// transaction manager postgres implementation
		fx.Provide(pgsqltx.New),
		// repository
		fx.Provide(
			sessionpgsql.New,
			userpgsql.New,
		),
		fx.Provide(
			paseto_tokens.New,
		),
		// application usecase
		fx.Provide(
			register.New,
			login.New,
		),
		// inject handlers group
		fx.Provide(
			fx.Annotate(
				post_login.New,
				fx.ResultTags(`group:"handlers"`),
			),
			fx.Annotate(
				post_register.New,
				fx.ResultTags(`group:"handlers"`),
			),
		),
		fx.Invoke(
			// annotations are struct tags alternative
			// we need to group handlers into slice and pass to api router
			fx.Annotate(httprest.New, fx.ParamTags(``, ``, `group:"handlers"`)),
		),
		fx.StartTimeout(time.Second),
	)

	app.Run()
}
