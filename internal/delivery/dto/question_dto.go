package dto

// CreateQuestionRequest create question payload
type CreateQuestion struct {
	Question        string            `json:"question"`
	Choose          map[string]string `json:"choose"`
	Answer          string            `json:"answer"`
	Description     string            `json:"description"`
	ImageUrl        string            `json:"image_url"`
	Question_Number int               `json:"question_number"`
	ExamID          uint              `json:"exam_id"`
}

// QuestionResponse response question payload

type QuestionResponse struct {
	ID              uint
	Question        string
	Choose          map[string]string
	Answer          string
	Description     string
	ImageUrl        string
	Question_Number int
	Exam            ExamResponse
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
