package model

import (
	"gorm.io/datatypes"
)

type QuestionModel struct {
	ID              uint `gorm:"primaryKey,autoIncrement"`
	Question        string
	Choose          datatypes.JSON `gorm:"type:jsonb"`
	Answer          string
	Question_number int
	Description     *string
	ImageURL        *string
	ExamID          uint
	Exam            ExamModel `gorm:"foreignKey:ExamID"`
}
