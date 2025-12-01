package exam

import (
	"examApp/internal/delivery/dto"
	do "examApp/internal/domain/entity"
	"examApp/internal/service/exam"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	service *exam.ExamService
}

func NewExamHandler(uc *exam.ExamService) *ExamHandler {
	return &ExamHandler{service: uc}
}

func (h *ExamHandler) CreateExam(c *gin.Context) {
	var dtos dto.CreateExam

	if err := c.ShouldBindJSON(&dtos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := do.NewExamKey(dtos.Subject, dtos.Departement, dtos.Year, dtos.Curriculum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	vals, er := h.service.ExamExistsByKey(*key)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Exam already exists with this value"})
		return
	}
	if vals {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Exam already exists with this value"})
		return
	}
	exam, er := do.NewExam(*key, dtos.AllowedTime)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": er.Error()})
		return
	}

	val, errs := h.service.CreateExam(exam)

	if errs != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dto.ExamResponse{
		ID:          val.ID(),
		Subject:     val.Subject(),
		Year:        val.Year(),
		Departement: val.Departement(),
		Curriculum:  val.Curriculum(),
		AllowedTime: val.AllowedTime(),
	}})
}

func (h *ExamHandler) CalulateOneScore(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	answer := c.Param("answer")
	ans, right, err := h.service.CalulateOneScore(uint(id), answer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is the answer correct": right, "Answer": ans})
}

func (h *ExamHandler) GetOneExam(c *gin.Context) {
	var q dto.ExamQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if q.Sort == "" {
		q.Sort = "question_number"
	}
	if q.SortOrder == "" {
		q.SortOrder = "asc"
	}
	if q.Page < 1 {
		q.Page = 1
	}
	if q.Limit < 1 {
		q.Limit = 2
	}

	out, total, err := h.service.GetOneExam(q.Sort, q.Subject, q.Departement, q.SortOrder, q.Page, q.Curriculum, q.Year, q.Limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := dto.MapExamDomaintoResponse(out)

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"page":           q.Page,
			"limit":          q.Limit,
			"total question": total,
		},
		"data": result,
	})
}

func (h *ExamHandler) GetAllExam(c *gin.Context) {
	out, err := h.service.GetAllExam()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	examresponse := make([]dto.ExamResponse, 0, len(out))
	for _, item := range out {
		examresponse = append(examresponse, dto.MapExamDomaintoResponse(&item))
	}

	c.JSON(http.StatusOK, gin.H{"data": examresponse})
}

func (h *ExamHandler) UpdateExam(c *gin.Context) {
	var dtos dto.CreateExam
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&dtos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := do.NewExamKey(dtos.Subject, dtos.Departement, dtos.Year, dtos.Curriculum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist, _ := h.service.ExamExists(uint(id))
	if exist == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "there is no exam at this id"})
		return
	}

	if check, _ := h.service.FindDuplicationForUpdate(*key, uint(id)); check {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Exam already exists with this value"})
		return
	}

	domain, errs := do.NewExam(*key, dtos.AllowedTime)

	if errs != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.Error()})
		return
	}

	val, er := h.service.UpdateExam(domain, uint(id))

	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated data": dto.ExamResponse{
		ID:          uint(id),
		Subject:     val.Subject(),
		Year:        val.Year(),
		Departement: val.Departement(),
		Curriculum:  val.Curriculum(),
		AllowedTime: val.AllowedTime(),
	}})
}

func (h *ExamHandler) CalulateAllScore(c *gin.Context) {
	var dtos dto.MockExam
	if err := c.ShouldBindJSON(&dtos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id []uint
	var answer []string
	for _, item := range dtos.AllAnswer {
		id = append(id, item.QuestionId)
		answer = append(answer, item.Answer)
	}
	total, err := h.service.CalulateAllScore(id, answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Total score": total})

}

func (h *ExamHandler) DeleteExam(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	existes, _ := h.service.ExamExists(uint(id))
	if existes == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "There is not any record at this id"})
		return
	}

	err := h.service.DeleteExam(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Deleted successfully"})
}
