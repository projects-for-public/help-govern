package models

import "time"

type Report struct {
	ID            int        `json:"id" gorm:"primaryKey"`
	Category      string     `json:"category" gorm:"not null"`
	Latitude      float64    `json:"latitude" gorm:"type:decimal(10,8);not null"`
	Longitude     float64    `json:"longitude" gorm:"type:decimal(11,8);not null"`
	Description   string     `json:"description" gorm:"type:text"`
	ReporterIP    string     `json:"-" gorm:"type:inet"`
	Status        string     `json:"status" gorm:"default:pending"`
	CreatedAt     time.Time  `json:"created_at"`
	VerifiedAt    *time.Time `json:"verified_at,omitempty"`
	StartedAt     *time.Time `json:"started_at,omitempty"`
	ResolvedAt    *time.Time `json:"resolved_at,omitempty"`
	ResolverNotes *string    `json:"resolver_notes,omitempty"`
	AdminNotes    *string    `json:"admin_notes,omitempty"`
	State         *string    `json:"state,omitempty"`
	City          *string    `json:"city,omitempty"`
	TwitterPosted bool       `json:"twitter_posted"`
	TwitterPostID *string    `json:"twitter_post_id,omitempty"`

	Images        []Image        `json:"images" gorm:"foreignKey:ReportID"`
	StatusUpdates []StatusUpdate `json:"timeline" gorm:"foreignKey:ReportID"`
}

func (Report) TableName() string {
	return "reports"
}

func (r *Report) GenerateShareURL() string {
	// TODO: Implement shareable URL generation
	return ""
}

func (r *Report) CanBeModifiedBy(userRole string) bool {
	// TODO: Implement permission logic
	return false
}
