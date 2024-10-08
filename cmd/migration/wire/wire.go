//go:build wireinject
// +build wireinject

package wire

import (
    "gin-init/internal/repository"
    "gin-init/internal/server"
    "gin-init/pkg/app"
    "gin-init/pkg/log"

    "github.com/google/wire"
    "github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
    repository.NewDB,
    repository.NewRedis,
    repository.NewElasticSearch,
    repository.NewRepository,
    repository.NewSettingRepository,
    repository.NewUserRepository,
    repository.NewPostRepository,
)
var serverSet = wire.NewSet(
    server.NewMigrate,
)

// build App
func newApp(migrate *server.Migrate) *app.App {
    return app.NewApp(
        app.WithServer(migrate),
        app.WithName("demo-migrate"),
    )
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
    panic(wire.Build(
        repositorySet,
        serverSet,
        newApp,
    ))
}
