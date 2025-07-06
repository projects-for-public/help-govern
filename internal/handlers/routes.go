package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers holds all handler dependencies for route registration.
type Handlers struct {
	Report *ReportHandler
	// Add other handlers here as needed, e.g. Auth *AuthHandler, Image *ImageHandler, etc.
}

// RegisterRoutes registers all application routes to the Gin engine.
// Pass in a Handlers struct with all required handlers.
func RegisterRoutes(r *gin.Engine, h *Handlers) {
	r.LoadHTMLGlob("web/templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Static("/static", "web/static")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/reports", h.Report.CreateReport)
	r.GET("/reports/:id", h.Report.GetReport)
	r.GET("/reports", h.Report.ListReports)
	r.PUT("/reports/:id", h.Report.UpdateReport)
	r.DELETE("/reports/:id", h.Report.DeleteReport)

	// Future: Add more routes for other handlers here
}
