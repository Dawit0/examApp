package question

import (
	"examApp/internal/delivery/dto"
	qus "examApp/internal/domain/entity"
	"examApp/internal/service/exam"
	"examApp/internal/service/question"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	service *question.QuestionService
	examuc  *exam.ExamService
}

func NewQuestionHandler(sv *question.QuestionService, esv *exam.ExamService) *QuestionHandler {
	return &QuestionHandler{service: sv, examuc: esv}
}

// Create Question godoc
// @Summary Create a new Question
// @Description add a new exam question
// @Tags Questions
// @Accept json
// @Produce json
// @Param request body dto.CreateQuestion true "Created Question "
// @Success 200 {object} dto.QuestionResponse "Question created successfully "
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /question/create [post]
func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	var model dto.CreateQuestion

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exams, err := h.examuc.ExamExists(model.ExamID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Exam with ID %d not found", model.ExamID)})
		return
	}

	examResponse := dto.ExamResponse{
		ID:          exams.ID(),
		Subject:     exams.Subject(),
		Year:        exams.Year(),
		Departement: exams.Departement(),
		Curriculum:  exams.Curriculum(),
		AllowedTime: exams.AllowedTime(),
	}

	val, errs := qus.NewQuestion(model.Question, model.Answer, model.Description, model.ImageUrl, model.Choose, model.Question_Number, model.ExamID)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}
	out, er := h.service.CreateQuestion(val)

	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dto.QuestionResponse{
		ID:              out.Id(),
		Question:        out.Question(),
		Choose:          out.Choose(),
		Answer:          out.Answer(),
		Question_Number: out.Questio_num(),
		Description:     out.Discription(),
		ImageUrl:        out.ImageUrl(),
		Exam:            examResponse,
	}})
}

// Get Question By ID
// @Summary Get Question by ID
// @Description use question id to get only one question
// @Tags Questions
// @Produce json
// @Param id path int true "Question ID"
// @Success 200 {object} dto.QuestionResponse "Question found"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /question/{id} [get]
func (h *QuestionHandler) GetOneQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	val, err := h.service.GetOneQuestion(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	domain := val.ReturnExam()

	c.JSON(http.StatusOK, gin.H{"data": dto.QuestionResponse{
		ID:              uint(id),
		Question:        val.Question(),
		Choose:          val.Choose(),
		Answer:          val.Answer(),
		Description:     val.Discription(),
		ImageUrl:        val.ImageUrl(),
		Question_Number: val.Questio_num(),
		Exam: dto.ExamResponse{
			ID:          domain.ID(),
			Subject:     domain.Subject(),
			Year:        domain.Year(),
			Departement: domain.Departement(),
			Curriculum:  domain.Curriculum(),
			AllowedTime: domain.AllowedTime(),
		},
	}})
}

// Get All Question godoc
// @Summary GetAll Question
// @Description All Question as a lis
// @Tags Questions
// @Produce json
// @Success 200 {object} []dto.QuestionResponse "ALL Question Founded"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /question [get]
func (h *QuestionHandler) GetAllQuestion(c *gin.Context) {
	val, err := h.service.GetAllQuestion()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	out := make([]dto.QuestionResponse, 0, len(val))

	for _, item := range val {
		domain := item.ReturnExam()
		out = append(out, dto.QuestionResponse{
			ID:              item.Id(),
			Question:        item.Question(),
			Choose:          item.Choose(),
			Answer:          item.Answer(),
			Description:     item.Discription(),
			ImageUrl:        item.ImageUrl(),
			Question_Number: item.Questio_num(),
			Exam: dto.ExamResponse{
				ID:          domain.ID(),
				Subject:     domain.Subject(),
				Year:        domain.Year(),
				Departement: domain.Departement(),
				Curriculum:  domain.Curriculum(),
				AllowedTime: domain.AllowedTime(),
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": out})
}

// Update Question godoc
// @Summary Update Question
// @Description Update Question by ID
// @Tags Questions
// @Accept json
// @Produce json
// @Param id path int true "Question ID"
// @Param request body dto.CreateQuestion true "Update Question"
// @Success 200 {object} dto.QuestionResponse "Question updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /question/{id} [put]
func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var model dto.CreateQuestion

	if errs := c.ShouldBindJSON(&model); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	exams, err := h.examuc.ExamExists(model.ExamID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Exam with ID %d not found", model.ExamID)})
		return
	}

	examResponse := dto.ExamResponse{
		ID:          exams.ID(),
		Subject:     exams.Subject(),
		Year:        exams.Year(),
		Departement: exams.Departement(),
		Curriculum:  exams.Curriculum(),
		AllowedTime: exams.AllowedTime(),
	}

	out, er := qus.NewQuestion(model.Question, model.Answer, "", "", model.Choose, model.Question_Number, model.ExamID)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	val, err := h.service.UpdateQuestion(uint(id), out)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated_Data": dto.QuestionResponse{
		ID:              uint(id),
		Question:        val.Question(),
		Choose:          val.Choose(),
		Answer:          val.Answer(),
		Description:     val.Discription(),
		ImageUrl:        val.ImageUrl(),
		Question_Number: val.Questio_num(),
		Exam:            examResponse,
	}})
}

// Delete Question godoc
// @Summary Delete Question
// @Description Delete Question by ID
// @Tags Questions
// @Produce json
// @Param id path int true "Question id"
// @Success 200 {object} map[string]string "Question Deleted Successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /question/{id} [delete]
func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.DeleteQuestion(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully"})
}
