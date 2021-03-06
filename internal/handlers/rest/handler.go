package rest

import (
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowMethods:           []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		AllowAllOrigins:        true,
		AllowBrowserExtensions: true,
	}))

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
		documents.GET("/get-client-contracts", h.getAllContractsByClient)
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
