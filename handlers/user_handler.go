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

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "successfully register user", Body: token})
}

func (h *compHandlers) LoginUserCredentials(c *gin.Context) {

	type Credentials struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var data Credentials

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: err.Error()})
		return
	}

	token, err := h.service.LoginUserCredentials(data.Email, data.Password)
	if err != nil {
		if err.Error() == "401" {
			c.JSON(http.StatusUnauthorized, dto.Response{Status: http.StatusUnauthorized, Error: "invalid email or password"})
			return
		} else if err.Error() == "404" {
			c.JSON(http.StatusNotFound, dto.Response{Status: http.StatusNotFound, Error: "email not found, please register"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "login successfully", Body: token})
}
