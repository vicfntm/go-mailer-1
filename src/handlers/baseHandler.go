package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vcfntm/go-mailer-1/src/services"
)

type Handler struct {
	services *services.Service
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	mailers := router.Group("/form")
	{
		mailers.POST("/feedback", handler.feedbackForm)
	}

	return router

}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}
