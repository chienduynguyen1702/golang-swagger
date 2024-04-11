package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListLectures godoc
// @Tags Lectures
// @Summary List all lectures
// @Description Get all lectures
// @Produce json
// @Success 200 {object} []Teacher
// @Router /api/v1/lectures [get]
func ListLectures(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all lectures successfully",
		"data":    lectures,
	})
}

// GetLecture godoc
// @Tags Lectures
// @Summary Get a lecture by ID
// @Description Get a lecture by ID
// @Produce json
// @Param id path int true "Lecture ID"
// @Success 200 {object} Teacher
// @Router /api/v1/lectures/{id} [get]
func GetLecture(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get a lecture successfully",
		"data":    lectures[intID],
	})
}

// CreateLecture godoc
// @Tags Lectures
// @Summary Create a lecture
// @Description Create a lecture
// @Accept json
// @Produce json
// @Param lecture body Teacher true "Lecture object"
// @Success 201 {object} Teacher
// @Router /api/v1/lectures [post]
func CreateLecture(c *gin.Context) {
	var lecture Teacher
	if err := c.ShouldBindJSON(&lecture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})
		return
	}
	lectures = append(lectures, lecture)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Lecture created successfully",
		"data":    lecture,
	})
}

// UpdateLecture godoc
// @Tags Lectures
// @Summary Update a lecture by ID
// @Description Update a lecture by ID
// @Accept json
// @Produce json
// @Param id path int true "Lecture ID"
// @Param lecture body Teacher true "Lecture object"
// @Success 200 {object} Teacher
// @Router /api/v1/lectures/{id} [put]
func UpdateLecture(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	var lecture Teacher
	if err := c.ShouldBindJSON(&lecture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})
		return
	}
	lectures[intID] = lecture
	c.JSON(http.StatusOK, gin.H{
		"message": "Lecture updated successfully",
		"data":    lecture,
	})
}

// DeleteLecture godoc
// @Tags Lectures
// @Summary Delete a lecture by ID
// @Description Delete a lecture by ID
// @Produce json
// @Param id path int true "Lecture ID"
// @Success 204
// @Router /api/v1/lectures/{id} [delete]
func DeleteLecture(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	lectures = append(lectures[:intID], lectures[intID+1:]...)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Lecture deleted successfully",
	})
}
