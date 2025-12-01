package entity

type Exam struct {
	id          uint
	key         ExamKey
	allowedTime int
	question    []Question
}

func NewExam(key ExamKey, allwed_time int) (*Exam, error) {

	return &Exam{
		key:         key,
		allowedTime: allwed_time,
	}, nil
}

// NewExamWithoutValidation creates an Exam without validation - use only when loading from database
func NewExamWithoutValidation(key ExamKey, allwed_time int) *Exam {
	return &Exam{
		key:         key,
		allowedTime: allwed_time,
	}
}

func (e Exam) ID() uint {
	return e.id
}

func (e Exam) ExamKey() ExamKey {
	return e.key
}

func (e Exam) Subject() string {
	return e.key.subject
}

func (e Exam) Year() int {
	return e.key.year
}

func (e Exam) Departement() string {
	return e.key.departement
}

func (e Exam) Curriculum() int {
	return e.key.curriculum
}

func (e Exam) AllowedTime() int {
	return e.allowedTime
}

func (e *Exam) SetID(id uint) {
	e.id = id
}

func (e Exam) Questions() []Question {
	return e.question
}

func (e *Exam) SetQuestion(qs []Question) {
	e.question = qs
}
