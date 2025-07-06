CREATE TABLE reports (
    id SERIAL PRIMARY KEY,
    category VARCHAR(50) NOT NULL,
    latitude DECIMAL(10, 8) NOT NULL,
    longitude DECIMAL(11, 8) NOT NULL,
    description TEXT,
    reporter_ip INET,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    verified_at TIMESTAMP,
    started_at TIMESTAMP,
    resolved_at TIMESTAMP,
    resolver_notes TEXT,
    admin_notes TEXT,
    state VARCHAR(50),
    city VARCHAR(100),
    twitter_posted BOOLEAN DEFAULT FALSE,
    twitter_post_id VARCHAR(50)
); 