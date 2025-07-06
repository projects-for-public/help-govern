-- Geographic queries
CREATE INDEX idx_reports_location ON reports(latitude, longitude);

-- Status filtering
CREATE INDEX idx_reports_status ON reports(status);

-- Time-based queries
CREATE INDEX idx_reports_created_at ON reports(created_at);

-- State-based filtering
CREATE INDEX idx_reports_state ON reports(state);

-- Image moderation
CREATE INDEX idx_images_moderation_status ON images(moderation_status); 