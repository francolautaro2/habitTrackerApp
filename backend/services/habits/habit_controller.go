package habits

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HabitController representa un controlador para gestionar operaciones relacionadas con hábitos.
type HabitController struct {
	HabitRepository HabitRepository
}

// CreateHabit es un controlador para crear un nuevo hábito.
func (controller *HabitController) CreateHabit(c *gin.Context) {
	var habit Habit
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

// GetHabit es un controlador para obtener un hábito por ID.
func (controller *HabitController) GetHabit(c *gin.Context) {
	id := c.Param("id")
	habit, err := controller.HabitRepository.GetHabit(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habit)
}

// DeleteHabit es un controlador para eliminar un hábito por ID.
func (controller *HabitController) DeleteHabit(c *gin.Context) {
	id := c.Param("id")
	err := controller.HabitRepository.DeleteHabit(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete habit", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted successfully"})
}

// UpdateHabit es un controlador para actualizar un hábito.
func (controller *HabitController) UpdateHabit(c *gin.Context) {
	var habit Habit
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

// GetAllHabits es un controlador para obtener todos los hábitos.
func (controller *HabitController) GetAllHabits(c *gin.Context) {
	habits, err := controller.HabitRepository.GetAllHabits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch habits", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habits)
}
