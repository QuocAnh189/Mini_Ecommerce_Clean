package http

import (
	_ "ecommerce_clean/docs"
	"ecommerce_clean/pkgs/middlewares"
	"ecommerce_clean/pkgs/minio"
	"ecommerce_clean/pkgs/token"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"ecommerce_clean/db"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/validation"
	"net/http"

	"ecommerce_clean/configs"
	"ecommerce_clean/pkgs/redis"

	cartHttp "ecommerce_clean/internals/cart/controller/http"
	orderHttp "ecommerce_clean/internals/order/controller/http"
	productHttp "ecommerce_clean/internals/product/controller/http"
	userHttp "ecommerce_clean/internals/user/controller/http"
)

type Server struct {
	engine      *gin.Engine
	cfg         *configs.Config
	validator   validation.Validation
	db          db.IDatabase
	minioClient *minio.MinioClient
	cache       redis.IRedis
	tokenMarker token.IMarker
	enforcer    *casbin.Enforcer
}

func NewServer(
	validator validation.Validation,
	db db.IDatabase,
	minioClient *minio.MinioClient,
	cache redis.IRedis,
	tokenMarker token.IMarker,
	enforcer *casbin.Enforcer,
) *Server {
	return &Server{
		engine:      gin.Default(),
		cfg:         configs.GetConfig(),
		validator:   validator,
		db:          db,
		minioClient: minioClient,
		cache:       cache,
		tokenMarker: tokenMarker,
		enforcer:    enforcer,
	}
}

func (s Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)
	if s.cfg.Environment == configs.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	s.engine.Use(func(c *gin.Context) {
		c.Set("enforcer", s.enforcer)
		c.Next()
	})

	s.engine.Use(middlewares.PrometheusMiddleware())
	s.engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s.engine.Use(middlewares.CorsMiddleware())

	if err := s.MapRoutes(); err != nil {
		logger.Fatalf("MapRoutes Error: %v", err)
	}

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Ecommerce Clean Architecture"})
	})

	//Start http server
	logger.Info("HTTP server is listening on PORT: ", s.cfg.HttpPort)
	if err := s.engine.Run(fmt.Sprintf(":%d", s.cfg.HttpPort)); err != nil {
		logger.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s Server) GetEngine() *gin.Engine {
	return s.engine
}

//	@title			Ecommerce Clean Architecture Swagger API
//	@version		1.0
// @host        	localhost:8080
// @BasePath    	/api/v1
//	@description	Swagger API for Go Clean Architecture.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Tran Phuoc Anh Quoc
//	@contact.email	anhquoc18092003@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func (s Server) MapRoutes() error {
	routesV1 := s.engine.Group("/api/v1")
	userHttp.Routes(routesV1, s.db, s.validator, s.minioClient, s.cache, s.tokenMarker)
	productHttp.Routes(routesV1, s.db, s.validator, s.minioClient, s.cache, s.tokenMarker)
	cartHttp.Routes(routesV1, s.db, s.validator, s.cache, s.tokenMarker)
	orderHttp.Routes(routesV1, s.db, s.validator, s.cache, s.tokenMarker)
	return nil
}
