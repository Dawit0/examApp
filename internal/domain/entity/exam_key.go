package entity

import "fmt"

type ExamKey struct {
	subject     string
	year        int
	departement string
	curriculum  int
}

func NewExamKey(subject, departement string, year, curriculum int) (*ExamKey, error) {
	if curriculum != 1 && curriculum != 2 {
		return nil, ErrInvalidCurriculum
	}

	if year < YearStart || year > YearLimit {
		return nil, ErrInvalidYear
	}

	if departement != "Natural" && departement != "Social" {
		return nil, ErrInvalidDepartement
	}

	return &ExamKey{
		subject:     subject,
		year:        year,
		departement: departement,
		curriculum:  curriculum,
	}, nil

}

func UpdatedExamKey(subject, departement string, curriculum, year int) *ExamKey {
	return &ExamKey{
		subject:     subject,
		departement: departement,
		year:        year,
		curriculum:  curriculum,
	}
}

func (e ExamKey) Subject() string {
	return e.subject
}

func (e ExamKey) Year() int {
	return e.year
}

func (e ExamKey) Departement() string {
	return e.departement
}

func (e ExamKey) Curriculum() int {
	return e.curriculum
}

func (k ExamKey) IsSame(other ExamKey) bool {
	return k.subject == other.subject && k.year == other.year && k.departement == other.departement && k.curriculum == other.curriculum
}

func (k ExamKey) String() string {
	return fmt.Sprintf("%s-%d-%s-%d", k.subject, k.year, k.departement, k.curriculum)
}
