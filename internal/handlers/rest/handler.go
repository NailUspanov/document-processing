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

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(204)
		//	return
		//}
		c.Next()
	})

	auth := router.Group("/auth")
	{
		//auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	documents := router.Group("/documents", h.userIdentity)
	{
		documents.POST("/create", h.create)
		documents.GET("/download/:filename", h.download)
		documents.GET("/child", h.getAllChildContracts)
		documents.GET("/adult", h.getAllAdultContracts)
	}

	clients := router.Group("/clients", h.userIdentity)
	{
		clients.GET("/", h.getAllClients)
	}

	courses := router.Group("/courses", h.userIdentity)
	{
		courses.POST("/create", h.createCourse)
		courses.GET("/", h.getAllCourses)
	}

	return router
}
