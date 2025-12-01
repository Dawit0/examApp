package main

import (
	dv1 "examApp/internal/delivery/exam"
	dv "examApp/internal/delivery/question"
	"examApp/internal/infrastructure/database"
	"examApp/internal/infrastructure/repository/exam"
	"examApp/internal/infrastructure/repository/question"
	"examApp/internal/infrastructure/router"
	"examApp/internal/pkg/logger"
	"examApp/internal/server/middleware"
	uc1 "examApp/internal/service/exam"
	uc "examApp/internal/service/question"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitZap("examApp")
	db := database.ConnectDB()

	logger.Log.Info("Database connected Successfully")

	repo := question.NewQuestionRepo(db)
	repo1 := exam.NewExamRepo(db)

	usecase := uc.NewQuestionService(repo)
	usecase1 := uc1.NewExamservice(repo1)

	handler := dv.NewQuestionHandler(usecase, usecase1)
	handler1 := dv1.NewExamHandler(usecase1)

	route := gin.New()

	route.Use(
		middleware.Recovery(),
		middleware.RequestLogger(),
	)

	router.QuestionRouter(handler, route)
	router.ExamRoute(handler1, route)

	logger.Log.Info("Server running on :8080")
	route.Run(":8080")
}
