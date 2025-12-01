package entity

type Question struct {
	id              uint
	question        string
	choose          map[string]string
	answer          string
	question_number int
	description     *string
	imageURL        *string
	examID          uint
	exam            Exam
}

func NewQuestion(question, answer, description, imageURL string, choose map[string]string, question_number int, examid uint) (*Question, error) {
	if len(answer) != 1 {
		return nil, ErrInvalidAnswer
	}

	if len(choose) != 4 {
		return nil, ErrInvalidChoose
	}

	return &Question{
		question:        question,
		choose:          choose,
		answer:          answer,
		question_number: question_number,
		description:     &description,
		imageURL:        &imageURL,
		examID:          examid,
	}, nil
}

func (q Question) Id() uint {
	return q.id
}

func (q Question) Question() string {
	return q.question
}

func (q Question) Choose() map[string]string {
	return q.choose
}

func (q Question) Answer() string {
	return q.answer
}

func (q Question) Questio_num() int {
	return q.question_number
}

func (q Question) Discription() string {
	return *q.description
}

func (q Question) ImageUrl() string {
	return *q.imageURL
}

func (q Question) ExamId() uint {
	return q.examID
}

func (q *Question) Set_Id(id uint) {
	q.id = id
}
func (q *Question) Exam(exam Exam) {
	q.exam = exam
}

func (q Question) ReturnExam() Exam {
	return q.exam
}

func (q Question) DescriptionPtr() *string {
	return q.description
}

func (q Question) ImageURLPtr() *string {
	return q.imageURL
}
