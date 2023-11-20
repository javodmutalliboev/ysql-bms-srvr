package administratorroutes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdministratorRoutes(subRouter *gin.RouterGroup) gin.HandlerFunc {
	return func(c *gin.Context) {
		administratorRouter := subRouter.Group("administrator")
		administratorRouter.GET("users", func(c *gin.Context) {
			c.String(http.StatusOK, "get users working")
		})
		c.Next()
	}
}
