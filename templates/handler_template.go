// templates/handler_template.go
// Template for HTTP handlers with proper error handling and logging

package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

// TODO: Implement handler for [FEATURE_NAME]
// Reference: feature-specification.md section [X]
// API Spec: api-documentation.md [ENDPOINT]
// Security: [AUTHENTICATION_REQUIREMENTS]
func [HandlerName](c *gin.Context) {
	// Step 1: Authentication/Authorization check
	// (if required per api-documentation.md)
	
	// Step 2: Rate limiting check
	// (as specified in api-documentation.md)
	
	// Step 3: Input validation
	// Validate according to API spec requirements
	
	// Step 4: Business logic
	// Call appropriate service methods
	
	// Step 5: Response formatting
	// Follow api-documentation.md response format
	
	// Step 6: Error handling
	// Use error codes from api-documentation.md
	
	// Step 7: Logging
	log.Printf("Handler [%s] executed for %s", "[HandlerName]", c.Request.URL.Path)
}

// Example: Reports listing handler template
func ListReportsTemplate(c *gin.Context) {
	// Implements GET /reports from api-documentation.md
	// Query params: lat, lng, zoom, category, status, limit
	// Returns: reports[] with clustering data
	
	// Extract query parameters
	lat := c.Query("lat")     // float
	lng := c.Query("lng")     // float  
	zoom := c.Query("zoom")   // int
	category := c.Query("category") // string
	status := c.Query("status")     // string
	limit := c.DefaultQuery("limit", "100") // int, default 100
	
	// TODO: Validate parameters
	// TODO: Call service to get reports with clustering
	// TODO: Return response as per API spec
}