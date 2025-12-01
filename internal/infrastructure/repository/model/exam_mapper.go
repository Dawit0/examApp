package model

import (
	"examApp/internal/domain/entity"
)

func MapDomaintoModels(e *entity.Exam) (*ExamModel, error) {
	return &ExamModel{
		ID:          e.ID(),
		Subject:     e.Subject(),
		Year:        e.Year(),
		Departement: e.Departement(),
		Curriculum:  e.Curriculum(),
		AllowedTime: e.AllowedTime(),
	}, nil
}

func MapModeltoDomains(m *ExamModel, skipValidation bool) (*entity.Exam, error) {
	var val *entity.Exam

	if skipValidation {
		// Use non-validating constructor for data already in database
		keys := entity.UpdatedExamKey(m.Subject, m.Departement, m.Curriculum, m.Year)
		val = entity.NewExamWithoutValidation(*keys, m.AllowedTime)
		val.SetID(m.ID)
	} else {
		// Use validating constructor for new data
		var err error
		key, errs := entity.NewExamKey(m.Subject, m.Departement, m.Year, m.Curriculum)
		if errs != nil {
			return nil, errs
		}
		val, err = entity.NewExam(*key, m.AllowedTime)
		if err != nil {
			return nil, err
		}
		val.SetID(m.ID)
	}

	domain := make([]entity.Question, 0)

	for _, item := range m.Questions {

		qst, errs := MapModeltoDomain(&item, true)
		if errs != nil {
			return nil, errs
		}

		domain = append(domain, *qst)

	}

	val.SetQuestion(domain)

	return val, nil
}
