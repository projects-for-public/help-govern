package models

type StatusUpdate struct {
	ID        int     `db:"id" json:"id"`
	ReportID  int     `db:"report_id" json:"report_id"`
	OldStatus *string `db:"old_status" json:"old_status,omitempty"`
	NewStatus string  `db:"new_status" json:"new_status"`
	Notes     *string `db:"notes" json:"notes,omitempty"`
	UpdatedBy *int    `db:"updated_by" json:"updated_by,omitempty"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}
