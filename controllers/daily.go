package controllers

import (
	"hitos_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDaily(c *gin.Context) {
	daily, err := models.GetDailys()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, daily)
}

func SetDaily(c *gin.Context) {
	var inJson models.Daily

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.SetDaily(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}
