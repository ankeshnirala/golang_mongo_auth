package routes

import (
	"github.com/ankeshnirala/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", controllers.Signup())
	incomingRoutes.POST("/signin", controllers.Login())
}