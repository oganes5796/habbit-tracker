package model

type Auth struct {
	ID        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type AuthInfo struct {
	Username string `json:"username"`
}
