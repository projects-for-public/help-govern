// templates/model_template.go
// Template for database models following database-schema.md

package models

import "time"

// TODO: Implement model for [TABLE_NAME]
// Reference: database-schema.md - [TABLE_NAME] table
// Fields: [LIST_ALL_FIELDS]
// Relationships: [LIST_RELATIONSHIPS]
type [ModelName] struct {
	// Primary key
	ID int `json:"id" gorm:"primaryKey"`
	
	// Required fields from schema
	// TODO: Add all fields from database-schema.md
	
	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	// TODO: Define associations as per schema
}

// TableName returns the table name for GORM
func ([ModelName]) TableName() string {
	return "[table_name]"
}

// TODO: Add model methods for business logic
// Reference feature requirements for what methods are needed

// Example: Report model template
type Report struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Category    string  `json:"category" gorm:"not null"`
	Latitude    float64 `json:"latitude" gorm:"type:decimal(10,8);not null"`
	Longitude   float64 `json:"longitude" gorm:"type:decimal(11,8);not null"`
	Description string  `json:"description" gorm:"type:text"`
	ReporterIP  string  `json:"-" gorm:"type:inet"` // Hidden from JSON
	Status      string  `json:"status" gorm:"default:pending"`
	
	// Timestamps from database-schema.md
	CreatedAt   time.Time  `json:"created_at"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`
	StartedAt   *time.Time `json:"started_at,omitempty"`  
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
	
	// Relationships
	Images        []Image        `json:"images" gorm:"foreignKey:ReportID"`
	StatusUpdates []StatusUpdate `json:"timeline" gorm:"foreignKey:ReportID"`
}

// Business methods based on feature requirements
func (r *Report) GenerateShareURL() string {
	// TODO: Implement shareable URL generation
	// Required by feature-specification.md - Anonymous Issue Reporting
	return ""
}

func (r *Report) CanBeModifiedBy(userRole string) bool {
	// TODO: Implement permission logic
	// Based on user roles from database-schema.md
	return false
}