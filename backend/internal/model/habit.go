package model

import "time"

type Habit struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	TargetValue *int      `json:"target_value,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type HabitInfo struct {
	UserID      int64  `json:"user_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	TargetValue *int   `json:"target_value,omitempty"`
}
