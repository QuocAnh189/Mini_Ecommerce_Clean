package initial

import (
	"ecommerce_clean/db"
	"ecommerce_clean/pkgs/logger"

	cartEntity "ecommerce_clean/internals/cart/entity"
	orderEntity "ecommerce_clean/internals/order/entity"
	productEntity "ecommerce_clean/internals/product/entity"
	userEntity "ecommerce_clean/internals/user/entity"
)

func InitMigrations(d *db.Database) error {
	if err := d.AutoMigrate(
		&userEntity.User{},
		&productEntity.Product{},
		&orderEntity.Order{},
		&orderEntity.OrderLine{},
		&cartEntity.Cart{},
		&cartEntity.CartLine{},
	); err != nil {
		logger.Fatal("Database migration failed", err)
		return err
	}

	logger.Info("Database migration completed successfully")
	return nil
}
