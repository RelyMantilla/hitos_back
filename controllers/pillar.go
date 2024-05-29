package controllers

import (
	"hitos_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPillar(c *gin.Context) {
	pillar, err := models.GetPillar()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, pillar)
}

func SetPillar(c *gin.Context) {

	var inJson models.Pillar

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.SetPillar(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": id})
}
