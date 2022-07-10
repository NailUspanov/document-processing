package rest

import (
	"github.com/gin-gonic/gin"
	"sstu/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	documents := router.Group("/documents")
	{
		documents.POST("/", h.create)
	}

	courses := router.Group("/courses")
	{
		courses.POST("/create", h.createCourse)
		courses.GET("/", h.getAll)
	}

	return router
}
