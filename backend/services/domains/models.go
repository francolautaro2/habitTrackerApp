package domains

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u UserClient) error
	GetUser(id string) (UserClient, error)
	DeleteUser(id string) error
	UpdateUser(u UserClient) error
	GetAllUsers() ([]UserClient, error)
	GetUserByUsernameOrEmail(usernameOrEmail string) (UserClient, error)
}

// Habit Interface
type HabitRepository interface {
	SaveHabit(habit Habit) (uint, error)
	GetHabit(id string) (Habit, error)
	UpdateHabit(habit Habit) error
	DeleteHabit(id string) error
	GetAllHabits() ([]Habit, error)
}

// Habit structure
type Habit struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type UserClient struct {
	gorm.Model
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Habits   []Habit `json:"habits" gorm:"foreignKey:UserID"` // Especifica la clave foránea aquí
}
