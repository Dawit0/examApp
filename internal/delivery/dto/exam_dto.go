package dto

type CreateExam struct {
	Subject     string `json:"subject" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	Departement string `json:"departement" binding:"required"`
	Curriculum  int    `json:"curriculum" binding:"required"`
	AllowedTime int    `json:"allowedTime" binding:"required"`
}

type ExamResponse struct {
	ID          uint
	Subject     string
	Year        int
	Departement string
	Curriculum  int
	AllowedTime int
	Question    []QuestionResponse
}

type ExamQuery struct {
	Sort        string `form:"sort_by"`
	SortOrder   string `form:"sort_order"`
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Subject     string `form:"subject"`
	Year        int    `form:"year"`
	Departement string `form:"departement"`
	Curriculum  int    `form:"curriculum"`
}

type ExamAnswers struct {
	QuestionId uint   `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
}

type MockExam struct {
	AllAnswer []ExamAnswers `json:"answer"`
}
