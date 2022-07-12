package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"io/ioutil"
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
	case "aptech_course": //TODO: make more meaningful name
		h.aptechCourse(c)
	case "registration_form_dksh":
		//TODO
	case "":
	default:
		logrus.Panicln("no such type in case")
	}

}

func (h *Handler) download(c *gin.Context) {

	type DownloadRequest struct {
		Filename string `uri:"filename"`
	}

	var input DownloadRequest
	if err := c.BindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	byteFile, err := ioutil.ReadFile("./data/documents/" + input.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+input.Filename)
	c.Data(http.StatusOK, "application/octet-stream", byteFile)
}

func (h *Handler) aptechCourse(c *gin.Context) {
	var input models.DocumentRequest
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	course, err := h.services.CourseService.GetByName(input.Course)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.DocumentService.Create(input, course)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:   false,
		Message: "document created",
		Data:    map[string]interface{}{"data": result},
	})

	//TODO: need to return doc (?)
}

func (h *Handler) getAllDocuments(c *gin.Context) {
	result, err := h.services.DocumentService.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:   false,
		Message: "documents found",
		Data:    map[string]interface{}{"data": result},
	})
}

func (h *Handler) getAllClients(c *gin.Context) {
	result, err := h.services.DocumentService.GetAllClients()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:   false,
		Message: "clients found",
		Data:    map[string]interface{}{"data": result},
	})
}
