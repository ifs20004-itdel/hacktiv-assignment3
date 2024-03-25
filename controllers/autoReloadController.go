package controllers

import (
	"assignment3/dto"
	"assignment3/models"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RenderStatus(c *gin.Context) {
	var result gin.H

	status := models.Status{}
	dto := dto.DTO{}
	data, err := os.ReadFile("status.json")
	if err != nil {
		result = gin.H{
			"message:": err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}
	err = json.Unmarshal(data, &status)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}
	water := status.Status["water"]
	wind := status.Status["wind"]

	if water < 5 {
		dto.WaterStatus = "Aman"
	} else if 6 <= water && water <= 8 {
		dto.WaterStatus = "Siaga"
	} else {
		dto.WaterStatus = "Bahaya"
	}

	if wind <= 6 {
		dto.WindStatus = "Aman"
	} else if 7 <= wind && wind <= 15 {
		dto.WindStatus = "Siaga"
	} else {
		dto.WindStatus = "Bahaya"
	}
	dto.Water = water
	dto.Wind = wind

	c.HTML(http.StatusOK, "index.html", dto)
}

func GetStatus(c *gin.Context) {
	var result gin.H

	status := models.Status{}
	data, err := os.ReadFile("status.json")
	if err != nil {
		result = gin.H{
			"message:": err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}
	err = json.Unmarshal(data, &status)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}
	c.JSON(http.StatusOK, status)
}
