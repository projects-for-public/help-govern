/ templates/service_template.go
// Template for service layer following feature-specification.md requirements

package services

import (
	"your-project/internal/models"
	"your-project/internal/database"
)

type [ServiceName]Service struct {
	db *database.DB
	// Add other dependencies (Cloudinary, Google Vision, etc.)
}

func New[ServiceName]Service(db *database.DB) *[ServiceName]Service {
	return &[ServiceName]Service{
		db: db,
	}
}

// TODO: Implement service methods based on feature requirements
// Reference: feature-specification.md - [FEATURE_NAME]
// Requirements: [LIST_KEY_REQUIREMENTS]
// Acceptance Criteria: [LIST_ACCEPTANCE_CRITERIA]

func (s *[ServiceName]Service) [MethodName]([params]) ([return_type], error) {
	// Step 1: Validate business rules
	// Step 2: Perform operations
	// Step 3: Handle errors appropriately
	// Step 4: Return results
	
	return nil, nil
}

// Example: Image processing service template
type ImageService struct {
	db          *database.DB
	cloudinary  *CloudinaryClient
	moderationAPI *GoogleVisionClient
}

// ProcessReportImages - Implements image handling from feature-specification.md
// Requirements: 
// - Upload to Cloudinary
// - Extract EXIF GPS data
// - Queue for moderation
// - Support JPEG, PNG, WebP
// - Max 5MB per image, 3 images per report
func (s *ImageService) ProcessReportImages(images []string, reportID int) error {
	// TODO: Implement following feature-specification.md requirements
	
	for _, imageData := range images {
		// 1. Validate image format and size
		// 2. Extract EXIF GPS coordinates
		// 3. Upload to Cloudinary
		// 4. Save metadata to database
		// 5. Queue for moderation
	}
	
	return nil
}
