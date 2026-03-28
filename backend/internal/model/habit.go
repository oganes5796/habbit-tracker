package model

import "time"

type Habit struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	TargetValue *int      `json:"target_value,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type HabitInfo struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	TargetValue *int   `json:"target_value,omitempty"`
}
