package habits

import (
	"habitTrackerApi/services/auth"
	"habitTrackerApi/services/domains"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	ALL HTTP LOGIC HABITS CONTROLLERS
*/

type HabitController struct {
	HabitRepository domains.HabitRepository
}

// Create habit
func (controller *HabitController) CreateHabit(c *gin.Context) {
	var habit domains.Habit
	if err := c.BindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := controller.HabitRepository.SaveHabit(habit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create habit", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// Get habit
func (controller *HabitController) GetHabit(c *gin.Context) {
	id := c.Param("id")
	habit, err := controller.HabitRepository.GetHabit(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habit)
}

// Delete habit
func (controller *HabitController) DeleteHabit(c *gin.Context) {
	id := c.Param("id")
	err := controller.HabitRepository.DeleteHabit(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete habit", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted successfully"})
}

// Update habit
func (controller *HabitController) UpdateHabit(c *gin.Context) {
	var habit domains.Habit
	if err := c.BindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := controller.HabitRepository.UpdateHabit(habit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update habit", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Habit updated successfully"})
}

// Get all habits
func (controller *HabitController) GetAllHabits(c *gin.Context) {

	// Extract ID from the token
	userID, err := auth.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get All habits from database
	habits, err := controller.HabitRepository.GetAllHabits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch habits", "details": err.Error()})
		return
	}

	// Filter all habits for user
	userHabits := []domains.Habit{}
	for _, habit := range habits {
		if habit.UserID == userID {
			userHabits = append(userHabits, habit)
		}
	}

	// Return all habits to the user
	c.JSON(http.StatusOK, userHabits)
}
