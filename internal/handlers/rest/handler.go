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
		documents.POST("/create", h.create)
		documents.GET("/download/:filename", h.download)
		documents.GET("/", h.getAllDocuments)
	}

	clients := router.Group("/clients")
	{
		clients.GET("/", h.getAllClients)
	}

	courses := router.Group("/courses")
	{
		courses.POST("/create", h.createCourse)
		courses.GET("/", h.getAllCourses)
	}

	return router
}
