# Civic Infrastructure Reporting Platform

## Project Overview

A web platform for citizens to report infrastructure issues (potholes, broken streetlights, water leaks, etc.) in India. Issues are displayed on an interactive map with clustering, moderation system, and automatic social media posting.

## Project Structure

```
help-govern/
├── README.md
├── docs/
│   ├── API.md
│   ├── DATABASE.md
│   ├── DEPLOYMENT.md
│   ├── FEATURES.md
│   └── FRONTEND.md
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── connection.go
│   │   └── migrations/
│   │       ├── 001_create_reports.sql
│   │       ├── 002_create_images.sql
│   │       ├── 003_create_users.sql
│   │       └── 004_create_status_updates.sql
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── reports.go
│   │   ├── images.go
│   │   ├── moderation.go
│   │   ├── admin.go
│   │   └── routes.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── ratelimit.go
│   │   └── cors.go
│   ├── models/
│   │   ├── report.go
│   │   ├── image.go
│   │   ├── user.go
│   │   └── status_update.go
│   ├── services/
│   │   ├── image_service.go
│   │   ├── moderation_service.go
│   │   ├── twitter_service.go
│   │   └── map_service.go
│   └── utils/
│       ├── exif.go
│       ├── clustering.go
│       └── validation.go
├── web/
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css
│   │   ├── js/
│   │   │   ├── map.js
│   │   │   ├── report.js
│   │   │   └── admin.js
│   │   └── images/
│   └── templates/
│       ├── base.html
│       ├── index.html
│       ├── admin.html
│       └── moderation.html
├── scripts/
│   ├── setup.sh
│   └── deploy.sh
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── .env.example
├── go.mod
└── go.sum
```

## Technology Stack

- **Backend**: Go with Gin framework
- **Database**: PostgreSQL (with SQLite option for development)
- **Frontend**: Vanilla JavaScript + HTML/CSS
- **Maps**: OpenStreetMap with Leaflet.js
- **Image Storage**: Cloudinary
- **Image Moderation**: Google Vision API / Azure Content Moderator
- **Social Media**: Twitter API v2
- **Deployment**: Docker ready, cloud hosting compatible

## Key Features for MVP

1. Anonymous issue reporting with categories
2. Interactive map with clustering
3. Photo upload with EXIF GPS extraction
4. Image moderation system
5. Admin/moderator roles and dashboard
6. Shareable issue URLs with timeline
7. Automatic Twitter posting with state authority tagging
8. Multi-language UI support
9. Rate limiting for anonymous users

## Development Phases

- **Week 1**: Core backend + map integration
- **Week 2**: Reporting system + image handling
- **Week 3**: User management + moderation
- **Week 4**: Social integration + polish
