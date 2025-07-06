package models

// TODO: move to JSON model definition as in report.go
type Category struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	NameHi      *string `db:"name_hi" json:"name_hi,omitempty"`
	Description *string `db:"description" json:"description,omitempty"`
	IconClass   *string `db:"icon_class" json:"icon_class,omitempty"`
	IsActive    bool    `db:"is_active" json:"is_active"`
	SortOrder   int     `db:"sort_order" json:"sort_order"`
}
