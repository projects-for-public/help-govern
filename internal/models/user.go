package models

// TODO: move to JSON model definition as in report.go
type User struct {
	ID           int     `db:"id" json:"id"`
	Username     string  `db:"username" json:"username"`
	Email        string  `db:"email" json:"email"`
	PasswordHash string  `db:"password_hash" json:"-"`
	Role         string  `db:"role" json:"role"`
	IsActive     bool    `db:"is_active" json:"is_active"`
	CreatedAt    string  `db:"created_at" json:"created_at"`
	LastLogin    *string `db:"last_login" json:"last_login,omitempty"`
	CreatedBy    *int    `db:"created_by" json:"created_by,omitempty"`
}
