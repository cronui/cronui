package config

import (
	"go.uber.org/fx"
)

func Module(configFile string) fx.Option {
	return fx.Options(
		fx.Provide(func() (*App, error) {
			return NewAppConfig(configFile)
		}),
		fx.Provide(provideDatabaseConfig),
		fx.Provide(provideServerConfig),
	)
}

func provideDatabaseConfig(app *App) *Database {
	return app.Database
}

func provideServerConfig(app *App) *Server {
	return app.Server
}
