package model

import (
	"encoding/json"
	"examApp/internal/domain/entity"
)

func MapDomaintoModel(q *entity.Question) (*QuestionModel, error) {
	val, err := json.Marshal(q.Choose())
	if err != nil {
		return nil, err
	}
	return &QuestionModel{
		ID:              q.Id(),
		Question:        q.Question(),
		Choose:          val,
		Answer:          q.Answer(),
		Question_number: q.Questio_num(),
		Description:     q.DescriptionPtr(),
		ImageURL:        q.ImageURLPtr(),
		ExamID:          q.ExamId(),
	}, nil
}

func MapModeltoDomain(m *QuestionModel, setquestion bool) (*entity.Question, error) {
	var val map[string]string
	err := json.Unmarshal(m.Choose, &val)
	if err != nil {
		return nil, err
	}

	// Handle nil pointers by providing empty strings as defaults
	description := ""
	if m.Description != nil {
		description = *m.Description
	}

	imageURL := ""
	if m.ImageURL != nil {
		imageURL = *m.ImageURL
	}

	out, errs := entity.NewQuestion(m.Question, m.Answer, description, imageURL, val, m.Question_number, m.ExamID)
	if errs != nil {
		return nil, errs
	}

	out.Set_Id(m.ID)

	if !setquestion {
		vals, err := MapModeltoDomains(&m.Exam, true)
		if err != nil {
			return nil, err
		}

		out.Exam(*vals)
	}

	return out, nil
}
