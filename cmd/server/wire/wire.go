//go:build wireinject
// +build wireinject

package wire

import (
    "gin-init/internal/handler"
    "gin-init/internal/repository"
    "gin-init/internal/server"
    "gin-init/internal/service"
    "gin-init/pkg/app"
    "gin-init/pkg/jwt"
    "gin-init/pkg/log"
    "gin-init/pkg/server/http"
    "gin-init/pkg/sid"

    "github.com/google/wire"
    "github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
    repository.NewDB,
    repository.NewRedis,
    repository.NewElasticSearch,
    repository.NewRepository,
    repository.NewTransaction,
    repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
    service.NewService,
    service.NewUserService,
)

var handlerSet = wire.NewSet(
    handler.NewHandler,
    handler.NewUserHandler,
)

var serverSet = wire.NewSet(
    server.NewHTTPServer,
    server.NewJob,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
    return app.NewApp(
        app.WithServer(httpServer, job),
        app.WithName("demo-server"),
    )
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
    panic(wire.Build(
        repositorySet,
        serviceSet,
        handlerSet,
        serverSet,
        sid.NewSid,
        jwt.NewJwt,
        newApp,
    ))
}
