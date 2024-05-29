package controllers

import (
	"hitos_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFamily(c *gin.Context) {
	family, err := models.GetFamily()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, family)
}

func SetFamily(c *gin.Context) {
	var inJson models.Family

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.SetFamily(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}
