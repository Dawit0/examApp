package question

import (
	"examApp/internal/domain/entity"
	model "examApp/internal/infrastructure/repository/model"

	"gorm.io/gorm"
)

// @Question Represent an exam question
// @Description Question model used for exam questions

type QuestionRepo struct {
	DB *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) *QuestionRepo {
	db.AutoMigrate(&model.QuestionModel{})
	return &QuestionRepo{DB: db}
}

func (rp *QuestionRepo) CreateQuestion(q *entity.Question) (*entity.Question, error) {

	models, err := model.MapDomaintoModel(q)

	if err != nil {
		return nil, err
	}

	errs := rp.DB.Create(&models).Error

	val, er := model.MapModeltoDomain(models, true)
	if er != nil {
		return nil, er
	}

	return val, errs

}

func (rp *QuestionRepo) GetOneQuestion(id uint) (*entity.Question, error) {
	var models model.QuestionModel

	err := rp.DB.Model(&model.QuestionModel{}).Preload("Exam").First(&models, id).Error

	result, errs := model.MapModeltoDomain(&models, false)

	if errs != nil {
		return nil, errs
	}

	return result, err
}

func (rp *QuestionRepo) GetAllQuestion() ([]entity.Question, error) {
	var models []model.QuestionModel

	err := rp.DB.Model(&model.QuestionModel{}).Preload("Exam").Find(&models).Error

	domain := make([]entity.Question, 0, len(models))

	for _, item := range models {
		val, errs := model.MapModeltoDomain(&item, false)
		if errs != nil {
			return nil, errs
		}

		domain = append(domain, *val)

	}

	return domain, err
}

func (rp *QuestionRepo) UpdateQuestion(id uint, q *entity.Question) (*entity.Question, error) {
	models, er := model.MapDomaintoModel(q)
	if er != nil {
		return nil, er
	}

	err := rp.DB.Model(&model.QuestionModel{}).Where("id=?", id).Updates(&models).Error

	domain, errs := model.MapModeltoDomain(models, true)

	if errs != nil {
		return nil, errs
	}

	return domain, err

}

func (rp *QuestionRepo) DeleteQuestion(id uint) error {

	var models model.QuestionModel
	err := rp.DB.Model(&model.QuestionModel{}).Where("id=?", id).Delete(&models).Error

	return err
}
