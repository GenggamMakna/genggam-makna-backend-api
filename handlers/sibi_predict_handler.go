package handlers

import (
	"genggam-makna-api/dto"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) SIBIImagePredict(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "image required"})
		return
	}
	defer file.Close()

	image_data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	cached, _ := h.service.GetPredictCache(image_data, dto.SIBI)
	if cached != nil {
		c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: cached, Message: "image predicted successfully"})
		return
	}

	result, err := h.service.SIBIImagePredict(image_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: result, Message: "image predicted successfully"})
}

func (h *compHandlers) SIBIVideoPredict(c *gin.Context) {
	file, _, err := c.Request.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "video required"})
		return
	}
	defer file.Close()

	video_data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	cached, _ := h.service.GetPredictCache(video_data, dto.SIBI)
	if cached != nil {
		c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: cached, Message: "video predicted successfully"})
		return
	}

	result, err := h.service.SIBIVideoPredict(video_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: result, Message: "video predicted successfully"})
}
