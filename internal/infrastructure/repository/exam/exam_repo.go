package exam

import (
	//"errors"
	"errors"
	"examApp/internal/domain/entity"
	"examApp/internal/infrastructure/repository/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// @Exam Represent an exam
// @Description Exam model used for exam questions

type ExamRepo struct {
	DB *gorm.DB
}

func NewExamRepo(db *gorm.DB) *ExamRepo {
	db.AutoMigrate(&model.ExamModel{})
	return &ExamRepo{DB: db}
}

func (h *ExamRepo) CreateExam(e *entity.Exam) (*entity.Exam, error) {
	models, er := model.MapDomaintoModels(e)
	if er != nil {
		return nil, er
	}
	err := h.DB.Create(&models).Error

	if err != nil {
		return nil, err
	}
	val, ers := model.MapModeltoDomains(models, true)
	if ers != nil {
		return nil, ers
	}
	return val, err

}

func (r ExamRepo) ExamExistsByKey(key entity.ExamKey) (bool, error) {
	var existe bool
	query := `
        SELECT EXISTS(
            SELECT 1 
            FROM exam_models 
            WHERE subject=$1 AND departement=$2 AND curriculum=$3 AND year=$4
        )
		`

	err := r.DB.Raw(query, key.Subject(), key.Departement(), key.Curriculum(), key.Year()).Scan(&existe).Error

	if err != nil {
		return false, err
	}
	return existe, nil
}

func (r ExamRepo) FindDuplicationForUpdate(key entity.ExamKey, excludeID uint) (bool, error) {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM exam_models
            WHERE subject = $1
              AND year = $2
              AND departement = $3
              AND curriculum = $4
              AND id <> $5  
        );
    `

	var exam bool
	var errs = r.DB.Raw(query, key.Subject(), key.Year(), key.Departement(), key.Curriculum(), excludeID).Scan(&exam).Error

	if errs != nil {

		return false, errs
	}

	return exam, nil
}

func (h *ExamRepo) GetOneExam(sortBy, subject, departement, sortOrder string,
	page, curriculum, year, limit int) (*entity.Exam, int64, error) {

	var models model.ExamModel
	if sortBy == "" {
		sortBy = "question_number"
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}

	var total int64

	offset := (page - 1) * limit

	err := h.DB.Model(&model.ExamModel{}).
		Where("subject = ?", subject).
		Where("departement = ?", departement).
		Where("curriculum = ?", curriculum).
		Where("year = ?", year).
		First(&models).Error

	if err != nil {
		return nil, 0, err
	}

	if err := h.DB.Model(&model.QuestionModel{}).Where("exam_id = ?", models.ID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page > int(total) {
		return nil, 0, errors.New("there is not any question at this page")
	}

	err = h.DB.Model(&model.ExamModel{}).Preload("Questions", func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).Limit(limit).Offset(offset)
	}).First(&models, models.ID).Error

	if err != nil {
		return nil, 0, err
	}

	if models.Questions == nil {
		models.Questions = []model.QuestionModel{}
	}

	val, errs := model.MapModeltoDomains(&models, true)

	if errs != nil {
		return nil, 0, errs
	}

	return val, total, nil
}

// ExamExists checks if an exam with the given ID exists
func (h *ExamRepo) ExamExists(id uint) (*entity.Exam, error) {

	var models model.ExamModel
	err := h.DB.Model(&model.ExamModel{}).First(&models, id).Error

	if err != nil {
		return nil, err
	}

	val, errs := model.MapModeltoDomains(&models, true)
	if errs != nil {
		return nil, errs
	}

	return val, nil
}

func (h *ExamRepo) GetAllExam() ([]entity.Exam, error) {

	var models []model.ExamModel

	err := h.DB.Model(&model.ExamModel{}).Preload("Questions").Find(&models).Error

	domain := make([]entity.Exam, 0, len(models))

	for _, item := range models {
		out, errs := model.MapModeltoDomains(&item, true)
		if errs != nil {
			continue
		}

		domain = append(domain, *out)
	}

	if err != nil {

		return nil, err
	}

	return domain, err
}

func (h *ExamRepo) CalulateOneScore(id uint, answer string) (correct_answer string, isright bool, err error) {
	var models model.QuestionModel

	err = h.DB.Model(&model.QuestionModel{}).First(&models, id).Error

	if err != nil {
		return "", false, err
	}

	if strings.EqualFold(models.Answer, answer) {
		return models.Answer, true, nil
	}

	return models.Answer, false, nil
}

func (h *ExamRepo) CalulateAllScore(id []uint, answer []string) (int64, error) {
	var models []model.QuestionModel

	err := h.DB.Model(&model.QuestionModel{}).Where("id IN ?", id).Find(&models).Error
	if err != nil {
		return 0, err
	}

	var total int64 = 0

	for i, item := range models {
		if strings.EqualFold(item.Answer, answer[i]) {
			total++
		}
	}
	return total, nil
}

func (h *ExamRepo) UpdateExam(id uint, e *entity.Exam) (*entity.Exam, error) {
	models, err := model.MapDomaintoModels(e)

	if err != nil {
		return nil, err
	}

	errs := h.DB.Model(&model.ExamModel{}).Where("id=?", id).Updates(&models).Error

	out, er := model.MapModeltoDomains(models, true)
	if er != nil {
		return nil, er
	}

	return out, errs
}

func (h *ExamRepo) DeleteExam(id uint) error {

	var models model.ExamModel
	err := h.DB.Model(&model.ExamModel{}).Where("id=?", id).Delete(&models).Error

	return err
}
