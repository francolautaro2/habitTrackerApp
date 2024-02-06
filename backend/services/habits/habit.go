package habits

import (
	"time"

	"gorm.io/gorm"
)

// Habit Interface
type HabitRepository interface {
	SaveHabit(habit Habit) (string, error)
	GetHabit(id string) (Habit, error)
	UpdateHabit(habit Habit) error
	DeleteHabit(id string) error
	GetAllHabits() ([]Habit, error)
}

// Habit structure
type Habit struct {
	gorm.Model
	UserID    uint
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// All habit controllers here
