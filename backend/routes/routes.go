package routes

import (
	"FullStack/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}

func MeetingRoutes(r *gin.Engine) {
	auth := r.Group("/meetings")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/", controllers.CreateMeeting)
		auth.GET("/", controllers.GetMeetings)
		auth.PUT("/:id", controllers.UpdateMeeting)
		auth.DELETE("/:id", controllers.DeleteMeeting)
		auth.POST("/:id/invite", controllers.InviteUser)
	}
}
