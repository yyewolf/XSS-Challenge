package api

import (
	"backend/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DelReport(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Check if id is valid
	iId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	// Delete report
	err = models.DelReport(iId)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func GetReport(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Check if id is valid
	iId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	// Get report
	report, err := models.GetReport(iId)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	}

	c.JSON(200, report)
}

func ListReports(c *gin.Context) {
	reports, err := models.GetAllReports()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, reports)
}

func NewReport(c *gin.Context) {
	// Get data
	var data models.Report
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	// Create report
	err = data.Create()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
