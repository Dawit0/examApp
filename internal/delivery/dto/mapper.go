package dto

import (
	"examApp/internal/domain/entity"
)

func MapExamDomaintoResponse(e *entity.Exam) ExamResponse {
	questions := make([]QuestionResponse, 0)
	for _, item := range e.Questions() {

		questions = append(questions, QuestionResponse{
			ID:              item.Id(),
			Question:        item.Question(),
			Choose:          item.Choose(),
			Answer:          item.Answer(),
			Question_Number: item.Questio_num(),
		})
	}

	return ExamResponse{
		ID:          e.ID(),
		Subject:     e.Subject(),
		Year:        e.Year(),
		Departement: e.Departement(),
		Curriculum:  e.Curriculum(),
		AllowedTime: e.AllowedTime(),
		Question:    questions,
	}
}
