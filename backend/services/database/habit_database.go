package database

import (
	"habitTrackerApi/services/domains"
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
func (r *DatabaseHabitRepository) SaveHabit(habit domains.Habit) (uint, error) {

	habit.CreatedAt = time.Now()
	habit.ExpiresAt = habit.CreatedAt.AddDate(0, 0, 30)

	result := r.Db.Create(&habit)
	if result.Error != nil {
		return 0, result.Error
	}
	return habit.UserID, nil
}

// Get a Habit
func (r *DatabaseHabitRepository) GetHabit(id string) (domains.Habit, error) {
	var habit domains.Habit
	result := r.Db.First(&habit, "id = ?", id)
	if result.Error != nil {
		return domains.Habit{}, result.Error
	}
	return habit, nil
}

// Update Habit
func (r *DatabaseHabitRepository) UpdateHabit(habit domains.Habit) error {
	result := r.Db.Save(&habit)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete Habit
func (r *DatabaseHabitRepository) DeleteHabit(id string) error {
	result := r.Db.Delete(&domains.Habit{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get all Habits from database
func (r *DatabaseHabitRepository) GetAllHabits() ([]domains.Habit, error) {
	var habits []domains.Habit
	result := r.Db.Find(&habits)
	if result.Error != nil {
		return nil, result.Error
	}
	return habits, nil
}
