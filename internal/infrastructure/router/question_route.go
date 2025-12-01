package router

import (
	"examApp/internal/delivery/question"

	"github.com/gin-gonic/gin"
)

func QuestionRouter(hands *question.QuestionHandler, r *gin.Engine) {
	api := r.Group("/api/v1/exam_app")
	{
		api.POST("/create", hands.CreateQuestion)
		api.GET("/:id", hands.GetOneQuestion)
		api.GET("", hands.GetAllQuestion)
		api.PUT("/:id", hands.UpdateQuestion)
		api.DELETE("/:id", hands.DeleteQuestion)
	}
}
