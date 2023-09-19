package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/harshau007/efl-backend/controllers"
	"github.com/harshau007/efl-backend/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}