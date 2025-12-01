package entity

import "errors"

var (
	ErrInvalidCurriculum  = errors.New("invalid curriculum type")
	ErrInvalidYear        = errors.New("invalid year type")
	ErrInvalidDepartement = errors.New("invalid departement")
)

const YearStart int = 2000
const YearLimit int = 2017

var Curriculem_list = [2]int{1, 2}
