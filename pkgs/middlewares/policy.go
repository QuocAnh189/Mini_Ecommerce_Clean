package middlewares

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthorizePolicy(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		role := roleVal.(string)
		e := c.MustGet("enforcer").(*casbin.Enforcer)
		ok, err := e.Enforce(role, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "enforce error"})
			c.Abort()
			return
		}
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
