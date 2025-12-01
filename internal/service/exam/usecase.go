package exam

import (
	do "examApp/internal/domain/entity"
	"examApp/internal/infrastructure/repository/exam"
)

type ExamService struct {
	repo *exam.ExamRepo
}

func NewExamservice(rp *exam.ExamRepo) *ExamService {
	return &ExamService{repo: rp}
}

func (uc *ExamService) CreateExam(e *do.Exam) (*do.Exam, error) {
	val, err := uc.repo.CreateExam(e)

	return val, err
}

func (uc *ExamService) ExamExistsByKey(key do.ExamKey) (bool, error) {
	return uc.repo.ExamExistsByKey(key)
}

func (uc *ExamService) FindDuplicationForUpdate(key do.ExamKey, excludeID uint) (bool, error) {
	return uc.repo.FindDuplicationForUpdate(key, excludeID)
}

func (uc *ExamService) GetOneExam(sortBy string, subject string, departement string, sortOrder string, page int, curriculum int, year int, limit int) (*do.Exam, int64, error) {

	val, total, err := uc.repo.GetOneExam(sortBy, subject, departement, sortOrder, page, curriculum, year, limit)

	return val, total, err
}

func (uc *ExamService) ExamExists(id uint) (*do.Exam, error) {
	val, err := uc.repo.ExamExists(id)
	return val, err
}

func (uc *ExamService) GetAllExam() ([]do.Exam, error) {
	val, err := uc.repo.GetAllExam()
	return val, err
}

func (uc *ExamService) UpdateExam(e *do.Exam, id uint) (*do.Exam, error) {
	val, err := uc.repo.UpdateExam(id, e)

	return val, err
}

func (uc *ExamService) CalulateOneScore(id uint, answer string) (string, bool, error) {
	return uc.repo.CalulateOneScore(id, answer)
}

func (uc *ExamService) CalulateAllScore(id []uint, answer []string) (int64, error) {
	return uc.repo.CalulateAllScore(id, answer)
}

func (uc *ExamService) DeleteExam(id uint) error {
	err := uc.repo.DeleteExam(id)
	return err
}
