package server

import (
	"gin-init/docs"
	"gin-init/internal/handler"
	"gin-init/internal/middleware"
	"gin-init/pkg/jwt"
	"gin-init/pkg/log"
	"gin-init/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	postHandler *handler.PostHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		// ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		// middleware.SignMiddleware(log),
	)

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			user := noAuthRouter.Group("/user")
			{
				user.POST("/register", userHandler.Register)
				user.POST("/login", userHandler.Login)
				user.POST("/bind/wechat", userHandler.BindWechat)
				user.POST("/post/list", postHandler.List)
				user.POST("/phone/send", userHandler.SendPhoneCode)
				user.POST("/phone/verify", userHandler.VerifyPhoneCode)
			}
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
			noAuthRouter.POST("/post/create", postHandler.Create)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
