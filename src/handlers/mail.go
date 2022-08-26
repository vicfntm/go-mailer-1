package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vcfntm/go-mailer-1/src/handlers/exceptionhandlers"
	"github.com/vcfntm/go-mailer-1/src/models"
)

var MESSAGE = "message did not sent"

func (handler *Handler) feedbackForm(c *gin.Context) {
	var input models.FeedbackForm

	if err := c.BindJSON(&input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if _, err := handler.services.Push(input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, MESSAGE)
	}

}
