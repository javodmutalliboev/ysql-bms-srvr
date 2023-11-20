package subroutes

import (
	"github.com/gin-gonic/gin"
	administratorroutes "ysql-bms/subroutes/administrator"
)

func SubRoutes(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		subRouter := router.Group("")
		subRouter.Use(administratorroutes.AdministratorRoutes(subRouter))
	}
}
