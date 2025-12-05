package main

import (
	"examApp/docs"
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

	_ "examApp/docs" // Import generated docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Exam API
// @version 1.0
// @description Exam API
// @host localhost:9090
// @BasePath /

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

	docs.SwaggerInfo.BasePath = "/"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Log.Info("Server running on :9090")
	route.Run(":9090")
}
