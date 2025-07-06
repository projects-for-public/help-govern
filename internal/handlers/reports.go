package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neerajkhandelwal/help-govern/internal/models"
	"github.com/neerajkhandelwal/help-govern/internal/services"
	"github.com/neerajkhandelwal/help-govern/internal/utils"
)

type ReportHandler struct {
	Service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{Service: service}
}

// ReportCreateRequest is the expected payload for report submission
// Only fields relevant to submission are included
// Validation tags are used for Gin binding
// Category is validated against the DB in the handler
// Images are ignored for now
//
type ReportCreateRequest struct {
	Category    string  `json:"category" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
	Description string  `json:"description"`
}

// POST /reports
func (h *ReportHandler) CreateReport(c *gin.Context) {
	var req ReportCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error("POST /reports - validation failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "VALIDATION_ERROR",
			"details": err.Error(),
		})
		return
	}
	// Validate latitude/longitude bounds
	if req.Latitude < -90 || req.Latitude > 90 || req.Longitude < -180 || req.Longitude > 180 {
		utils.Error("POST /reports - invalid coordinates: lat=%v, lng=%v", req.Latitude, req.Longitude)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "VALIDATION_ERROR",
			"details": "Invalid latitude or longitude.",
		})
		return
	}
	// Validate category exists
	exists, err := h.Service.CategoryExists(c.Request.Context(), req.Category)
	if err != nil {
		utils.Error("POST /reports - failed to check category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_ERROR",
			"details": "Could not validate category.",
		})
		return
	}
	if !exists {
		utils.Error("POST /reports - category not found: %s", req.Category)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "VALIDATION_ERROR",
			"details": "Invalid category.",
		})
		return
	}
	// Create the report
	report := models.Report{
		Category:    req.Category,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Description: req.Description,
		Status:      "pending",
	}
	if err := h.Service.CreateReport(c.Request.Context(), &report); err != nil {
		utils.Error("POST /reports - failed to create report: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_ERROR",
			"details": err.Error(),
		})
		return
	}
	shareURL := report.GenerateShareURL()
	c.JSON(http.StatusCreated, gin.H{
		"id":        report.ID,
		"share_url": shareURL,
		"message":   "Report submitted successfully",
	})
}

// GET /reports/:id
func (h *ReportHandler) GetReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error("GET /reports/:id - invalid report ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return
	}
	report, err := h.Service.GetReportByID(c.Request.Context(), id)
	if err != nil {
		utils.Error("GET /reports/:id - failed to fetch report: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch report", "details": err.Error()})
		return
	}
	if report == nil {
		utils.Info("GET /reports/:id - report not found (id=%d)", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}
	c.JSON(http.StatusOK, report)
}

// GET /reports
func (h *ReportHandler) ListReports(c *gin.Context) {
	reports, err := h.Service.ListReports(c.Request.Context())
	if err != nil {
		utils.Error("GET /reports - failed to list reports: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list reports", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
}

// PUT /reports/:id
func (h *ReportHandler) UpdateReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error("PUT /reports/:id - invalid report ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return
	}
	var req models.Report
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error("PUT /reports/:id - validation failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	req.ID = id
	if err := h.Service.UpdateReport(c.Request.Context(), &req); err != nil {
		utils.Error("PUT /reports/:id - failed to update report: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update report", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// DELETE /reports/:id
func (h *ReportHandler) DeleteReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error("DELETE /reports/:id - invalid report ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return
	}
	if err := h.Service.DeleteReport(c.Request.Context(), id); err != nil {
		utils.Error("DELETE /reports/:id - failed to delete report: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete report", "details": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
