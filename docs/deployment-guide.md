# Deployment Guide

## Environment Setup

### Development Environment

#### Prerequisites

```bash
# Install Go 1.21+
curl -OL https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Install PostgreSQL
sudo apt update
sudo apt install postgresql postgresql-contrib

# Install Node.js for frontend tools (optional)
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs
```

#### Database Setup

```bash
# Create database and user
sudo -u postgres psql
CREATE DATABASE help_govern;
CREATE USER help_govern_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE help_govern TO help_govern_user;
\q

# Run migrations
cd help-govern
go run cmd/migrate/main.go
```

#### Environment Variables

```bash
# Create .env file
cp .env.example .env

# Edit .env with your values
DATABASE_URL=postgres://help_govern_user:your_password@localhost/help_govern
CLOUDINARY_URL=cloudinary://api_key:api_secret@cloud_name
GOOGLE_VISION_API_KEY=your_google_vision_key
TWITTER_BEARER_TOKEN=your_twitter_bearer_token
JWT_SECRET=your_jwt_secret
SERVER_PORT=8080
```

### Production Deployment Options

## Option 1: Cloud Hosting (Recommended)

### Railway Deployment

```bash
# Install Railway CLI
npm install -g @railway/cli

# Login and deploy
railway login
railway init
railway add postgres
railway deploy

# Set environment variables in Railway dashboard
```

### Render Deployment

```bash
# Connect GitHub repository to Render
# Set build command: go build -o main cmd/server/main.go
# Set start command: ./main
# Add environment variables in Render dashboard
```

### Vercel Deployment (Serverless)

```bash
# Install Vercel CLI
npm install -g vercel

# Deploy
vercel

# Configure serverless functions in vercel.json
```

## Option 2: Raspberry Pi Deployment

### System Requirements

- Raspberry Pi 4 with 4GB+ RAM
- 32GB+ SD card (Class 10)
- Stable internet connection
- External storage for images (optional)

### Pi Setup

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
sudo usermod -aG docker pi

# Install Docker Compose
sudo apt install docker-compose

# Setup swap file for better performance
sudo dphys-swapfile swapoff
sudo nano /etc/dphys-swapfile
# Set CONF_SWAPSIZE=2048
sudo dphys-swapfile setup
sudo dphys-swapfile swapon
```

### Docker Deployment

```yaml
# docker-compose.yml
version: '3.8'

services:
  web:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://help_govern_user:password@db:5432/help_govern
      - CLOUDINARY_URL=${CLOUDINARY_URL}
      - GOOGLE_VISION_API_KEY=${GOOGLE_VISION_API_KEY}
      - TWITTER_BEARER_TOKEN=${TWITTER_BEARER_TOKEN}
      - JWT_SECRET=${JWT_SECRET}
    depends_on:
      - db
    restart: unless-stopped

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_DB=help_govern
      - POSTGRES_USER=help_govern_user
      - POSTGRES_PASSWORD=password
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    restart: unless-stopped

  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - web
    restart: unless-stopped

volumes:
  postgres_data:
```

### SSL Setup with Let's Encrypt

```bash
# Install Certbot
sudo apt install certbot python3-certbot-nginx

# Get SSL certificate
sudo certbot --nginx -d yourdomain.com

# Auto-renewal
sudo crontab -e
# Add: 0 12 * * * /usr/bin/certbot renew --quiet
```

## Option 3: Traditional VPS Deployment

### Server Setup (Ubuntu 22.04)

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install dependencies
sudo apt install -y git nginx postgresql postgresql-contrib certbot python3-certbot-nginx

# Create application user
sudo adduser --system --group civic
sudo mkdir -p /var/www/help-govern
sudo chown civic:civic /var/www/help-govern
```

### Application Deployment

```bash
# Clone repository
sudo -u civic git clone https://github.com/projects-for-prople/help-govern.git /var/www/help-govern
cd /var/www/help-govern

# Build application
sudo -u civic go build -o help-govern-server cmd/server/main.go

# Create systemd service
sudo nano /etc/systemd/system/help-govern.service
```

```ini
[Unit]
Description=Help Govern Server
After=network.target

[Service]
Type=simple
User=civic
Group=civic
WorkingDirectory=/var/www/help-govern
ExecStart=/var/www/help-govern/help-govern-server
Restart=always
RestartSec=5
Environment=DATABASE_URL=postgres://help_govern_user:password@localhost/help_govern
Environment=PORT=8080

[Install]
WantedBy=multi-user.target
```

```bash
# Enable and start service
sudo systemctl enable help-govern
sudo systemctl start help-govern
sudo systemctl status help-govern
```

### Nginx Configuration

```nginx
# /etc/nginx/sites-available/help-govern
server {
    listen 80;
    server_name yourdomain.com www.yourdomain.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name yourdomain.com www.yourdomain.com;

    ssl_certificate /etc/letsencrypt/live/yourdomain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/yourdomain.com/privkey.pem;

    # Security headers
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header Referrer-Policy "no-referrer-when-downgrade" always;
    add_header Content-Security-Policy "default-src 'self' http: https: data: blob: 'unsafe-inline'" always;

    # Static files
    location /static/ {
        alias /var/www/help-govern/web/static/;
        expires 1y;
        add_header Cache-Control "public, immutable";
    }

    # API routes
    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }

    # Main application
    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }
}
```

## Database Migration and Backup

### Migration Scripts

```bash
# Create migration script
#!/bin/bash
# scripts/migrate.sh
set -e

echo "Running database migrations..."

# Connect to database and run migrations
psql $DATABASE_URL -f internal/database/migrations/001_create_reports.sql
psql $DATABASE_URL -f internal/database/migrations/002_create_images.sql
psql $DATABASE_URL -f internal/database/migrations/003_create_users.sql
psql $DATABASE_URL -f internal/database/migrations/004_create_status_updates.sql

echo "Migrations completed successfully!"
```

### Backup Strategy

```bash
#!/bin/bash
# scripts/backup.sh
set -e

BACKUP_DIR="/var/backups/help-govern"
DATE=$(date +%Y%m%d_%H%M%S)

# Create backup directory
mkdir -p $BACKUP_DIR

# Database backup
pg_dump $DATABASE_URL > $BACKUP_DIR/database_$DATE.sql

# Compress backup
gzip $BACKUP_DIR/database_$DATE.sql

# Keep only last 30 days of backups
find $BACKUP_DIR -name "database_*.sql.gz" -mtime +30 -delete

echo "Backup completed: $BACKUP_DIR/database_$DATE.sql.gz"
```

### Automated Backups

```bash
# Add to crontab
crontab -e

# Daily backup at 2 AM
0 2 * * * /var/www/help-govern/scripts/backup.sh

# Weekly cleanup
0 3 * * 0 /var/www/help-govern/scripts/cleanup.sh
```

## Monitoring and Logging

### Application Logging

```bash
# Create log directory
sudo mkdir -p /var/log/help-govern
sudo chown civic:civic /var/log/help-govern

# Configure log rotation
sudo nano /etc/logrotate.d/help-govern
```

```
/var/log/help-govern/*.log {
    daily
    missingok
    rotate 30
    compress
    delaycompress
    notifempty
    create 644 civic civic
    postrotate
        systemctl reload help-govern
    endscript
}
```

### System Monitoring

```bash
# Install monitoring tools
sudo apt install htop iotop nethogs

# Simple monitoring script
#!/bin/bash
# scripts/monitor.sh
while true; do
    echo "=== $(date) ==="
    echo "CPU Usage:"
    top -bn1 | grep "Cpu(s)"
    echo "Memory Usage:"
    free -h
    echo "Disk Usage:"
    df -h /
    echo "Active Connections:"
    netstat -ant | grep :8080 | wc -l
    echo "========================"
    sleep 300
done
```

## Security Considerations

### Firewall Setup

```bash
# Configure UFW
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow ssh
sudo ufw allow 'Nginx Full'
sudo ufw enable
sudo ufw status
```

### Security Headers

```nginx
# Add to nginx config
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
add_header X-Frame-Options "SAMEORIGIN" always;
add_header X-Content-Type-Options "nosniff" always;
add_header X-XSS-Protection "1; mode=block" always;
```

### Regular Updates

```bash
#!/bin/bash
# scripts/update.sh
set -e

echo "Stopping application..."
sudo systemctl stop help-govern

echo "Pulling latest code..."
sudo -u civic git pull origin main

echo "Building application..."
sudo -u civic go build -o help-govern-server cmd/server/main.go

echo "Running migrations..."
./scripts/migrate.sh

echo "Starting application..."
sudo systemctl start help-govern

echo "Update completed!"
```

## Performance Optimization

### Database Optimization

```sql
-- Add indexes for common queries
CREATE INDEX CONCURRENTLY idx_reports_location ON reports USING GIST (ST_Point(longitude, latitude));
CREATE INDEX CONCURRENTLY idx_reports_status_created ON reports(status, created_at);
CREATE INDEX CONCURRENTLY idx_images_moderation ON images(moderation_status);

-- Analyze tables
ANALYZE reports;
ANALYZE images;
ANALYZE users;
```

### Caching Strategy

```nginx
# Static file caching
location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg)$ {
    expires 1y;
    add_header Cache-Control "public, immutable";
}

# API response caching
location /api/reports {
    proxy_cache reports_cache;
    proxy_cache_valid 200 5m;
    proxy_pass http://localhost:8080;
}
```

## Troubleshooting

### Common Issues

1. **Database Connection Issues**
   - Check DATABASE_URL format
   - Verify PostgreSQL service status
   - Check firewall rules

2. **Image Upload Failures**
   - Verify Cloudinary credentials
   - Check file size limits
   - Monitor disk space

3. **High Memory Usage**
   - Monitor Go garbage collection
   - Check for memory leaks
   - Optimize database queries

4. **Slow Map Loading**
   - Implement proper clustering
   - Add database indexes
   - Use CDN for static assets

### Log Analysis

```bash
# View application logs
sudo journalctl -u help-govern -f

# View nginx logs
sudo tail -f /var/log/nginx/access.log
sudo tail -f /var/log/nginx/error.log

# View PostgreSQL logs
sudo tail -f /var/log/postgresql/postgresql-15-main.log
```