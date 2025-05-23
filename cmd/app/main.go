package main

import (
	"ecommerce_clean/configs"
	"ecommerce_clean/db"
	"ecommerce_clean/pkgs/casbin"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/mail"
	"ecommerce_clean/pkgs/minio"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
	"ecommerce_clean/pkgs/validation"
	"sync"

	cartEntity "ecommerce_clean/internals/cart/entity"
	orderEntity "ecommerce_clean/internals/order/entity"
	productEntity "ecommerce_clean/internals/product/entity"
	httpServer "ecommerce_clean/internals/server/http"
	userEntity "ecommerce_clean/internals/user/entity"
)

var wg sync.WaitGroup

func main() {
	cfg := configs.LoadConfig()
	logger.Initialize(cfg.Environment)

	database, err := db.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	enforcer, err := casbin.InitCasbinEnforcer(database)
	if err != nil {
		logger.Fatal(err)
	}

	err = database.AutoMigrate(
		&userEntity.User{},
		&productEntity.Product{},
		&orderEntity.Order{},
		&orderEntity.OrderLine{},
		&cartEntity.Cart{},
		&cartEntity.CartLine{},
	)
	if err != nil {
		logger.Fatal("Database migration fail", err)
	}

	validator := validation.New()

	//minio
	minioClient, err := minio.NewMinioClient(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioBucket,
		cfg.MinioBaseurl,
		cfg.MinioUseSSL,
	)
	if err != nil {
		logger.Fatalf("Failed to connect to MinIO: %s", err)
	}

	//mailer
	mailer := mail.NewMailer(
		cfg.MailHost,
		cfg.MailPort,
		cfg.MailUser,
		cfg.MailPassword,
		cfg.MailFrom,
	)

	//cache
	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	//token
	tokenMaker, err := token.NewJTWMarker()
	if err != nil {
		logger.Fatal(err)
	}

	httpSvr := httpServer.NewServer(validator, database, minioClient, cache, tokenMaker, mailer, enforcer)

	wg.Add(1)

	// Run HTTP server
	go func() {
		defer wg.Done()
		if err := httpSvr.Run(); err != nil {
			logger.Fatal("Running HTTP server error:", err)
		}
	}()

	wg.Wait()
}
