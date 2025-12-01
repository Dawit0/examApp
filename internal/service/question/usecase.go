package question

import (
	qu "examApp/internal/domain/entity"
	"examApp/internal/infrastructure/repository/question"
)

type QuestionService struct {
	repo *question.QuestionRepo
}

func NewQuestionService(rp *question.QuestionRepo) *QuestionService {
	return &QuestionService{repo: rp}
}

func (uc *QuestionService) CreateQuestion(q *qu.Question) (*qu.Question, error) {
	val, err := uc.repo.CreateQuestion(q)

	return val, err
}

func (uc *QuestionService) GetOneQuestion(id uint) (*qu.Question, error) {
	val, err := uc.repo.GetOneQuestion(id)
	return val, err
}

func (uc *QuestionService) GetAllQuestion() ([]qu.Question, error) {
	val, err := uc.repo.GetAllQuestion()

	return val, err
}

func (uc *QuestionService) UpdateQuestion(id uint, q *qu.Question) (*qu.Question, error) {
	val, err := uc.repo.UpdateQuestion(id, q)

	return val, err
}

func (uc *QuestionService) DeleteQuestion(id uint) error {
	err := uc.repo.DeleteQuestion(id)

	return err
}
