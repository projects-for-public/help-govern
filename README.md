# Help Govern

> Empowering citizens to report infrastructure issues through transparent, map-based reporting.

<!-- [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) -->
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-green.svg)](https://www.postgresql.org/)

## 🚀 Overview

A web platform that enables citizens across India to anonymously report infrastructure issues like potholes, broken streetlights, water leaks, and poor drainage. Reports are displayed on an interactive map with intelligent clustering, go through a moderation system, and are automatically shared on social media with relevant state authorities tagged.

### Key Features

- **Anonymous Reporting**: No account required - citizens can report issues instantly
- **Interactive Map**: OpenStreetMap integration with smart clustering based on zoom level
- **Photo Upload**: EXIF GPS extraction and automated image moderation
- **Status Tracking**: Real-time updates from submission to resolution
- **Social Integration**: Automatic Twitter posting with state authority mentions
- **Multi-language**: Hindi and English support with extensible i18n framework
- **Admin Dashboard**: Comprehensive moderation and management interface

## 🌟 Use Cases

### For Citizens
- Report potholes, streetlight issues, water leaks instantly
- Track report status from submission to resolution
- Share issues with community via social media
- No registration required - completely anonymous

### For Authorities
- Real-time visibility into citizen complaints
- Geographic clustering to identify problem areas
- Automated social media notifications
- Transparent status tracking and accountability

### For Civic Organizations
- Data-driven advocacy with issue statistics
- Evidence collection through photos and GPS
- Community engagement through shared reporting
- Government response monitoring


## 📋 Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 15+
- Node.js 18+ (for frontend tooling)

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/civic-report.git
cd civic-report

# Copy environment variables
cp .env.example .env

# Install dependencies
go mod download

# Setup database
createdb civic_reports
go run cmd/migrate/main.go

# Run the application
go run cmd/server/main.go
```

The application will be available at `http://localhost:8080`

## 🏗️ Architecture

### Technology Stack

- **Backend**: Go + Gin framework
- **Database**: PostgreSQL with spatial indexing
- **Frontend**: Vanilla JavaScript + Leaflet.js
- **Image Storage**: Cloudinary
- **Maps**: OpenStreetMap
- **Social Media**: Twitter API v2

### Project Structure

See [project-structure](docs/project-structure.md) for the structuring of the files under the project.

## 📖 Documentation

Our comprehensive documentation covers all aspects of the project:

- **[Feature Specifications](docs/feature-specification.md)** - Detailed requirements and user stories for all features
- **[API Documentation](docs/api-documentation.md)** - Complete API reference with endpoints and examples
- **[Database Schema](docs/database-schema.md)** - Database design, relationships, and sample data
- **[Frontend Specifications](docs/frontend-specification.md)** - UI components, styling guidelines, and interactions
- **[Deployment Guide](docs/deployment-guide.md)** - Production deployment options and configurations

## 🛠️ Development

### Environment Setup

1. **Database Configuration**
   ```bash
   # Create database and user
   sudo -u postgres psql
   CREATE DATABASE civic_reports;
   CREATE USER civic_user WITH PASSWORD 'your_password';
   GRANT ALL PRIVILEGES ON DATABASE civic_reports TO civic_user;
   ```

2. **Environment Variables**
   ```bash
   # Edit .env file with your configurations
   DATABASE_URL=postgres://civic_user:password@localhost/civic_reports
   CLOUDINARY_URL=cloudinary://api_key:api_secret@cloud_name
   GOOGLE_VISION_API_KEY=your_google_vision_key
   TWITTER_BEARER_TOKEN=your_twitter_bearer_token
   JWT_SECRET=your_jwt_secret
   ```

3. **Run Development Server**
   ```bash
   # Backend
   go run cmd/server/main.go
   
   # Frontend (auto-reloads on changes)
   # No build process required - vanilla JS/HTML/CSS
   ```

### Key Components

- **Report Submission**: Anonymous reporting with GPS detection and photo upload
- **Map Clustering**: Dynamic marker clustering based on zoom level and geographic proximity
- **Moderation System**: Image approval workflow with automated content screening
- **Admin Interface**: User management, report status updates, and analytics
- **Social Integration**: Automated Twitter posting with state authority tagging

## 📊 Development Roadmap

Follow [TASKS.md](TASKS.md) and [feature-specifications.md](docs/feature-specification.md) docs.

## 🤝 Contributing

We welcome contributions! Here's how to get started:

1. **Fork the Repository**
   ```bash
   git fork https://github.com/yourusername/civic-report.git
   ```

2. **Create Feature Branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```

3. **Make Changes**
   - Follow Go conventions and add tests
   - Update documentation if needed
   - Ensure all tests pass

4. **Submit Pull Request**
   - Describe changes clearly
   - Link to related issues
   - Ensure CI passes

### Development Guidelines

- **Code Style**: Follow Go conventions with `gofmt` and `golint`
- **Testing**: Write unit tests for new features
- **Documentation**: Update docs for API changes
- **Security**: Never commit sensitive data or credentials


## 📱 Screenshots

*Coming soon - screenshots of the map interface, reporting form, and admin dashboard*

## 🚀 Deployment

The platform supports multiple deployment options:

- **Cloud Hosting**: Railway, Render, or Vercel for serverless deployment
- **Traditional VPS**: Ubuntu/CentOS with Nginx reverse proxy
- **Raspberry Pi**: Low-cost local deployment for communities
- **Docker**: Containerized deployment with docker-compose

See the [Deployment Guide](docs/deployment-guide.md) for detailed instructions and configurations.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **OpenStreetMap** for providing free, open geographic data
- **Cloudinary** for image storage and optimization
- **Font Awesome** for iconography
- **The Go Community** for excellent libraries and tools

## 📞 Support

- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/yourusername/civic-report/issues)
- 💡 **Feature Requests**: [GitHub Discussions](https://github.com/yourusername/civic-report/discussions)
- 📧 **Email**: contact@civic-report.org
- 🐦 **Twitter**: [@CivicReportIndia](https://twitter.com/CivicReportIndia)

---

**Made with ❤️ for Indian citizens by the open source community**