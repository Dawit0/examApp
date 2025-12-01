package router

import (
	"examApp/internal/delivery/exam"

	"github.com/gin-gonic/gin"
)

func ExamRoute(hands *exam.ExamHandler, r *gin.Engine) {
	api := r.Group("/exam")
	{
		api.POST("", hands.CreateExam)
		api.GET("/", hands.GetOneExam)
		api.GET("", hands.GetAllExam)
		api.PUT("/:id", hands.UpdateExam)
		api.GET("/score/:id/:answer", hands.CalulateOneScore)
		api.POST("/score", hands.CalulateAllScore)
		api.DELETE("/:id", hands.DeleteExam)
	}
}
