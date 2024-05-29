package controllers

import (
	"hitos_back/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPersonByID(c *gin.Context) {
	in := c.Query("ID")
	id, err := strconv.Atoi(in)
	//models.DB.Where("name like ?", "%"+inJson.Name+"%").Find(&person)
	person, err := models.GetPersonId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, person)
}

func GetPerson(c *gin.Context) {
	// nombre := c.Query("nombre")
	var inJson models.InPerson

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//models.DB.Where("name like ?", "%"+inJson.Name+"%").Find(&person)
	person, err := models.GetPerson(inJson.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, person)
}

func SetPerson(c *gin.Context) {
	var inJson models.Pillar

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//result := models.DB.Create(&inJson)
	res, err := models.SetPerson(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
