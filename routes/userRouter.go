package routes

import (
	"github.com/ankeshnirala/golang-jwt/controllers"
	"github.com/ankeshnirala/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.Use(middleware.Authenticate())
	
	incomingRoutes.GET("/users",  controllers.AllUsers())
	incomingRoutes.GET("users/:id", controllers.AUser())
}