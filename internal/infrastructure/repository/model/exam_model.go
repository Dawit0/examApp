package model

type ExamModel struct {
	ID          uint   `gorm:"primeryKey,autoIncrement"`
	Subject     string `gorm:"index"`
	Year        int    `gorm:"index"`
	Departement string `gorm:"index"`
	Curriculum  int    `gorm:"index"`
	AllowedTime int
	Questions   []QuestionModel `gorm:"foreignKey:ExamID" json:"-"`
}
