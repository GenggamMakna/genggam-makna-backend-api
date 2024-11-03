package handlers

import (
	"genggam-makna-api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterUserCredential(c *gin.Context) {
	var data dto.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: err.Error()})
		return
	}

	if data.Password == "" {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "password can't be null"})
		return
	}

	token, err := h.service.RegisterUserCredential(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "Successfully register user", Body: token})
}
