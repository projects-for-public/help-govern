package services

import (
	"context"
	"errors"

	"github.com/projects-for-public/help-govern/internal/models"
	"gorm.io/gorm"
)

type ReportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{db: db}
}

// CategoryExists checks if a category exists in the database
// TODO: Add caching
func (s *ReportService) CategoryExists(ctx context.Context, name string) (bool, error) {
	var count int64
	err := s.db.WithContext(ctx).Model(&models.Category{}).Where("name = ? AND is_active = TRUE", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateReport creates a new report
func (s *ReportService) CreateReport(ctx context.Context, report *models.Report) error {
	return s.db.WithContext(ctx).Create(report).Error
}

// GetReportByID fetches a report by its ID
func (s *ReportService) GetReportByID(ctx context.Context, id int) (*models.Report, error) {
	var report models.Report
	err := s.db.WithContext(ctx).Preload("Images").Preload("StatusUpdates").First(&report, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &report, err
}

// ListReports returns all reports (optionally with filters)
func (s *ReportService) ListReports(ctx context.Context) ([]models.Report, error) {
	var reports []models.Report
	err := s.db.WithContext(ctx).Preload("Images").Preload("StatusUpdates").Find(&reports).Error
	return reports, err
}

// UpdateReport updates an existing report
func (s *ReportService) UpdateReport(ctx context.Context, report *models.Report) error {
	return s.db.WithContext(ctx).Save(report).Error
}

// DeleteReport deletes a report by ID
func (s *ReportService) DeleteReport(ctx context.Context, id int) error {
	return s.db.WithContext(ctx).Delete(&models.Report{}, id).Error
}
