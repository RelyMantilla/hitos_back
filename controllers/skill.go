package controllers

import (
	"hitos_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSkill(c *gin.Context) {
	skill, err := models.GetSkill()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, skill)
}

func SetSkill(c *gin.Context) {
	var inJson models.Skill

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.SetSkill(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
