package controllers

import (
	"hitos_back/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCompetence(c *gin.Context) {

	pillarid := c.Query("pillarid")
	id, _ := strconv.Atoi(pillarid)
	competence, err := models.GetCompetence(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, competence)
}

func SetCompetence(c *gin.Context) {
	var inJson models.Competence

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//result := models.DB.Create(&inJson)
	id, err := models.SetCompetence(inJson)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, id)
}
