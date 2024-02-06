package database

import (
	"habitTrackerApi/services/habits"
	"time"

	"gorm.io/gorm"
)

type DatabaseHabitRepository struct {
	Db *gorm.DB
}

// Create a new database habit repository
func NewDatabaseHabitRepository(db *gorm.DB) *DatabaseHabitRepository {
	return &DatabaseHabitRepository{
		Db: db,
	}
}

// create habit in database
func (r *DatabaseHabitRepository) SaveHabit(habit habits.Habit) (string, error) {

	habit.CreatedAt = time.Now()
	habit.ExpiresAt = habit.CreatedAt.AddDate(0, 0, 30)

	result := r.Db.Create(&habit)
	if result.Error != nil {
		return "", result.Error
	}
	return habit.ID, nil
}

// Get a Habit
func (r *DatabaseHabitRepository) GetHabit(id string) (habits.Habit, error) {
	var habit habits.Habit
	result := r.Db.First(&habit, "id = ?", id)
	if result.Error != nil {
		return habits.Habit{}, result.Error
	}
	return habit, nil
}

// Update Habit
func (r *DatabaseHabitRepository) UpdateHabit(habit habits.Habit) error {
	result := r.Db.Save(&habit)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete Habit
func (r *DatabaseHabitRepository) DeleteHabit(id string) error {
	result := r.Db.Delete(&habits.Habit{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get all Habits from database
func (r *DatabaseHabitRepository) GetAllHabits() ([]habits.Habit, error) {
	var habits []habits.Habit
	result := r.Db.Find(&habits)
	if result.Error != nil {
		return nil, result.Error
	}
	return habits, nil
}
