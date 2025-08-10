package main

import (
	"ecommerce_clean/configs"
	"ecommerce_clean/db"
	"ecommerce_clean/pkgs/casbin"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
	"ecommerce_clean/pkgs/validation"
	"sync"

	"ecommerce_clean/internals/initial"
	httpServer "ecommerce_clean/internals/server/http"
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

	err = initial.InitMigrations(database)
	if err != nil {
		logger.Fatal("Database migration failed", err)
	}

	validator := validation.New()

	//minio
	minioClient, err := initial.InitMinioClient(
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
	mailer := initial.InitMailer(cfg.MailHost, cfg.MailPort, cfg.MailUser, cfg.MailPassword, cfg.MailFrom)

	// twilio
	twilioProvider, err := initial.InitTwilioProvider(cfg.TwilioAccountSID, cfg.TwilioAuthToken, cfg.TwilioFromNumber, cfg.TwilioServiceID)
	if err != nil {
		logger.Fatalf("Failed to initialize Twilio provider: %s", err)
	}

	//redis
	redis, err := initial.InitRedis(redis.Config{Address: cfg.RedisURI, Password: cfg.RedisPassword, Database: cfg.RedisDB})
	if err != nil {
		logger.Fatalf("Failed to initialize Redis: %s", err)
	}

	//token
	tokenMaker, err := token.NewJTWMarker()
	if err != nil {
		logger.Fatal(err)
	}

	httpSvr := httpServer.NewServer(validator, database, minioClient, redis, tokenMaker, mailer, twilioProvider, enforcer)

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
