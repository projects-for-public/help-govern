CREATE TABLE images (
    id SERIAL PRIMARY KEY,
    report_id INTEGER REFERENCES reports(id) ON DELETE CASCADE,
    cloudinary_url VARCHAR(255) NOT NULL,
    cloudinary_public_id VARCHAR(255) NOT NULL,
    image_type VARCHAR(20) DEFAULT 'report', -- 'report' or 'resolution'
    moderation_status VARCHAR(20) DEFAULT 'pending',
    moderation_notes TEXT,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    moderated_at TIMESTAMP,
    moderated_by INTEGER REFERENCES users(id)
); 