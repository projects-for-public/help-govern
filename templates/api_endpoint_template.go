// templates/api_endpoint_template.go
// Template for API endpoints following api-documentation.md specifications

package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"your-project/internal/models"
	"your-project/internal/services"
)

// TODO: Implement as per api-documentation.md
// Endpoint: [ENDPOINT_PATH]
// Method: [HTTP_METHOD] 
// Authentication: [PUBLIC/PROTECTED]
// Rate Limit: [RATE_LIMIT_SPEC]
func [HandlerName](c *gin.Context) {
	// Step 1: Validate input according to API spec
	// Required fields: [LIST_REQUIRED_FIELDS]
	// Optional fields: [LIST_OPTIONAL_FIELDS]
	
	// Step 2: Extract and validate request data
	var request [RequestStruct]
	if err := c.ShouldBindJSON(&request); err != nil {
		// Return validation error as per api-documentation.md error format
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed",
			"code": "VALIDATION_ERROR",
			"details": err.Error(),
		})
		return
	}
	
	// Step 3: Business logic implementation
	// Call appropriate service method
	// Handle service-specific errors
	
	// Step 4: Return response as per API specification
	// Success response format from api-documentation.md
	c.JSON(http.StatusOK, gin.H{
		// Response fields as specified in docs
	})
}

// Example: POST /reports endpoint template
func CreateReportTemplate(c *gin.Context) {
	// Implementation following api-documentation.md POST /reports spec
	// Required: category, latitude, longitude, description, images[]
	// Returns: id, share_url, message
	
	var request struct {
		Category    string   `json:"category" binding:"required"`
		Latitude    float64  `json:"latitude" binding:"required"`
		Longitude   float64  `json:"longitude" binding:"required"`
		Description string   `json:"description"`
		Images      []string `json:"images"` // base64 encoded
	}
	
	// TODO: Implement validation, processing, and response
}