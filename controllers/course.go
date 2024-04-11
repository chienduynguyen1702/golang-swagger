package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListCourses godoc
// @Tags Courses
// @Summary List all courses
// @Description Get all courses
// @Produce json
// @Success 200 {object} []Course
// @Router /api/v1/courses [get]
func ListCourses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all courses successfully",
		"data":    courses,
	})
}

// GetCourse godoc
// @Tags Courses
// @Summary Get a course by ID
// @Description Get a course by ID
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} Course
// @Router /api/v1/courses/{id} [get]
func GetCourse(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get a course successfully",
		"data":    courses[intID],
	})
}

// DeleteCourse godoc
// @Tags Courses
// @Summary Delete a course by ID
// @Description Delete a course by ID
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} Course
// @Router /api/v1/courses/{id} [delete]
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a course successfully",
		"data":    courses[intID],
	})
}
