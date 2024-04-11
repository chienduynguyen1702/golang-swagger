package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListStudents godoc
// @Tags Students
// @Summary List all students
// @Description Get all students
// @Produce json
// @Success 200 {object} []Student
// @Router /api/v1/students [get]
func ListStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all students successfully",
		"data":    students,
	})
}

// GetStudent godoc
// @Tags Students
// @Summary Get a student by ID
// @Description Get a student by ID
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} Student
// @Router /api/v1/students/{id} [get]
func GetStudent(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get a student successfully",
		"data":    students[intID],
	})
}

// CreateStudent godoc
// @Tags Students
// @Summary Create a student
// @Description Create a student
// @Accept json
// @Produce json
// @Param student body Student true "Student object"
// @Success 201 {object} Student
// @Router /api/v1/students [post]
func CreateStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})
		return
	}
	students = append(students, student)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Create a student successfully",
		"data":    student,
	})
}

// UpdateStudent godoc
// @Tags Students
// @Summary Update a student by ID
// @Description Update a student by ID
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body Student true "Student object"
// @Success 200 {object} Student
// @Router /api/v1/students/{id} [put]
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})
		return
	}
	students[intID] = student
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a student successfully",
		"data":    student,
	})
}

// DeleteStudent godoc
// @Tags Students
// @Summary Delete a student by ID
// @Description Delete a student by ID
// @Param id path int true "Student ID"
// @Success 204
// @Router /api/v1/students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	students = append(students[:intID], students[intID+1:]...)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Delete a student successfully",
	})
}
