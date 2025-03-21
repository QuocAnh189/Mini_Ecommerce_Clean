package response

import (
	"github.com/gin-gonic/gin"

	"ecommerce_clean/configs"
)

type ErrorResponse struct {
	Data interface{} `json:"data"`
}

func Error(c *gin.Context, status int, err error, message string) {
	cfg := configs.GetConfig()
	errorRes := map[string]interface{}{
		"message": message,
	}

	if cfg.Environment != configs.ProductionEnv {
		errorRes["debug"] = err.Error()
	}

	c.JSON(status, ErrorResponse{Data: errorRes})
}
