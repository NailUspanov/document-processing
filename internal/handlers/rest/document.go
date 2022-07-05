package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
	"sstu/internal/models"
)

type Type struct {
	Doctype string `json:"doctype"`
}

func (h *Handler) create(c *gin.Context) {
	var doctype Type
	if err := c.ShouldBindBodyWith(&doctype, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	switch doctype.Doctype {
	case "java_retraining_declaration":
		h.javaTrainingDeclaration(c)
	case "registration_form_dksh":
		//TODO
	case "":
	default:
		logrus.Panicln("no such type in case")
	}

}

func (h *Handler) javaTrainingDeclaration(c *gin.Context) {
	var input models.DocumentRequest
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.DocumentService.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:   false,
		Message: "document created",
		Data:    map[string]interface{}{"data": result},
	})
}
