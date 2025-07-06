# Database Schema

## Tables Overview

### 1. Reports Table

Stores all infrastructure issue reports.

```sql
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
```

**Status Values**: pending, verified, in_progress, resolved, rejected

### 2. Images Table

Stores image metadata for reports and resolutions.

```sql
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
```

**Moderation Status Values**: pending, approved, rejected

### 3. Users Table

Stores moderator and admin user accounts.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'moderator',
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    created_by INTEGER REFERENCES users(id)
);
```

**Role Values**: moderator, admin

### 4. Status Updates Table

Tracks timeline of status changes for reports.

```sql
CREATE TABLE status_updates (
    id SERIAL PRIMARY KEY,
    report_id INTEGER REFERENCES reports(id) ON DELETE CASCADE,
    old_status VARCHAR(20),
    new_status VARCHAR(20) NOT NULL,
    notes TEXT,
    updated_by INTEGER REFERENCES users(id),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 5. Categories Table

Predefined categories for infrastructure issues.

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    name_hi VARCHAR(50), -- Hindi translation
    description TEXT,
    icon_class VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE,
    sort_order INTEGER DEFAULT 0
);
```

**Initial Categories**:

- potholes
- broken_streetlight
- no_streetlight
- water_leaks
- poor_drainage
- damaged_sidewalk
- accident_prone
- garbage_heap
- wrong_side_driving

### 6. State Authorities Table

Twitter handles of state authorities for auto-posting.

```sql
CREATE TABLE state_authorities (
    id SERIAL PRIMARY KEY,
    state VARCHAR(50) NOT NULL,
    authority_name VARCHAR(100) NOT NULL,
    twitter_handle VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE
);
```

## Indexes for Performance

```sql
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
```

## Sample Data Insertion

```sql
-- Insert categories
INSERT INTO categories (name, name_hi, description, icon_class) VALUES
('potholes', 'गड्ढे', 'Road potholes and surface damage', 'fas fa-road'),
('broken_streetlight', 'टूटी स्ट्रीट लाइट', 'Non-functioning street lights', 'fas fa-lightbulb'),
('water_leaks', 'पानी का रिसाव', 'Water pipe leaks and wastage', 'fas fa-tint'),
('poor_drainage', 'खराब जल निकासी', 'Waterlogging and drainage issues', 'fas fa-water'),
('damaged_sidewalk', 'क्षतिग्रस्त फुटपाथ', 'Broken or damaged sidewalks', 'fas fa-walking');

-- Insert sample state authorities
INSERT INTO state_authorities (state, authority_name, twitter_handle) VALUES
('Rajasthan', 'Rajasthan Police', '@RajasthanPolice'),
('Delhi', 'Delhi Government', '@AAPDelhi'),
('Maharashtra', 'Mumbai Police', '@MumbaiPolice');
```
