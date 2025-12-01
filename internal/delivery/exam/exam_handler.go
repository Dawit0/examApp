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

// CreateExam godoc
// @Summary Create a new Exam
// @Description Create a new exam with subject, department, year, curriculum and allowed time
// @Tags Exam
// @Accept json
// @Produce json
// @Param request body dto.CreateExam true "Create Exam"
// @Success 200 {object} dto.ExamResponse "Exam created successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /exam [post]
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

// CalulateOneScore godoc
// @Summary Calculate score for one question
// @Description Calculate if the answer for a specific question is correct
// @Tags Exam
// @Produce json
// @Param id path int true "Question ID"
// @Param answer path string true "Answer"
// @Success 200 {object} map[string]interface{} "Score calculated"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /exam/score/{id}/{answer} [get]
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

// GetOneExam godoc
// @Summary Get exam with questions
// @Description Get exam details with paginated questions based on query parameters
// @Tags Exam
// @Produce json
// @Param subject query string false "Subject"
// @Param departement query string false "Department"
// @Param year query int false "Year"
// @Param curriculum query int false "Curriculum"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(2)
// @Param sort query string false "Sort field" default(question_number)
// @Param sortOrder query string false "Sort order (asc/desc)" default(asc)
// @Success 200 {object} dto.ExamResponse "Exam found"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /exam/ [get]
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

// GetAllExam godoc
// @Summary Get all exams
// @Description Get list of all exams
// @Tags Exam
// @Produce json
// @Success 200 {array} dto.ExamResponse "All exams"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /exam [get]
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

// UpdateExam godoc
// @Summary Update exam
// @Description Update exam details by ID
// @Tags Exam
// @Accept json
// @Produce json
// @Param id path int true "Exam ID"
// @Param request body dto.CreateExam true "Update Exam"
// @Success 200 {object} dto.ExamResponse "Exam updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /exam/{id} [put]
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

// CalulateAllScore godoc
// @Summary Calculate total score for multiple questions
// @Description Calculate total score based on multiple question answers
// @Tags Exam
// @Accept json
// @Produce json
// @Param request body dto.MockExam true "Mock exam answers"
// @Success 200 {object} map[string]interface{} "Total score"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /exam/score [post]
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

// DeleteExam godoc
// @Summary Delete exam
// @Description Delete exam by ID
// @Tags Exam
// @Produce json
// @Param id path int true "Exam ID"
// @Success 200 {object} map[string]string "Exam deleted successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /exam/{id} [delete]
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
