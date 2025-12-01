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

func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.DeleteQuestion(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully"})
}
