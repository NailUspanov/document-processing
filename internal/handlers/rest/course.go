package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"sstu/internal/models"
)

func (h *Handler) createCourse(c *gin.Context) {
	var input models.CourseRequest

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.CourseService.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:   false,
		Message: "course created",
		Data:    map[string]interface{}{"data": result},
	})

}

func (h *Handler) getAllCourses(c *gin.Context) {
	result, err := h.services.CourseService.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:   false,
		Message: "courses found",
		Data:    map[string]interface{}{"data": result},
	})

}

func (h *Handler) getCourseByName(c *gin.Context) {
	name := c.Query("name")
	result, err := h.services.CourseService.GetByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:   false,
		Message: "course found",
		Data:    map[string]interface{}{"data": result},
	})

}
