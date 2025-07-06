package models

// TODO: move to JSON model definition as in report.go
type Image struct {
	ID                 int     `db:"id" json:"id"`
	ReportID           int     `db:"report_id" json:"report_id"`
	CloudinaryURL      string  `db:"cloudinary_url" json:"cloudinary_url"`
	CloudinaryPublicID string  `db:"cloudinary_public_id" json:"cloudinary_public_id"`
	ImageType          string  `db:"image_type" json:"image_type"`
	ModerationStatus   string  `db:"moderation_status" json:"moderation_status"`
	ModerationNotes    *string `db:"moderation_notes" json:"moderation_notes,omitempty"`
	UploadedAt         string  `db:"uploaded_at" json:"uploaded_at"`
	ModeratedAt        *string `db:"moderated_at" json:"moderated_at,omitempty"`
	ModeratedBy        *int    `db:"moderated_by" json:"moderated_by,omitempty"`
}
