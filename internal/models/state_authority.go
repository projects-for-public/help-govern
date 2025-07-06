package models

// TODO: move to JSON model definition as in report.go
type StateAuthority struct {
	ID            int     `db:"id" json:"id"`
	State         string  `db:"state" json:"state"`
	AuthorityName string  `db:"authority_name" json:"authority_name"`
	TwitterHandle *string `db:"twitter_handle" json:"twitter_handle,omitempty"`
	IsActive      bool    `db:"is_active" json:"is_active"`
}
