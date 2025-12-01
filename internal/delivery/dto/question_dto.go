package dto

type CreateQuestion struct {
	Question        string            `json:"question"`
	Choose          map[string]string `json:"choose"`
	Answer          string            `json:"answer"`
	Description     string            `json:"description"`
	ImageUrl        string            `json:"image_url"`
	Question_Number int               `json:"question_number"`
	ExamID          uint              `json:"exam_id"`
}

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
