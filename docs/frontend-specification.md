# Frontend Specifications

## Technology Stack

- **Framework**: Vanilla JavaScript (no framework dependencies)
- **Maps**: Leaflet.js with OpenStreetMap
- **UI Library**: Custom CSS with modern design principles
- **Icons**: Font Awesome 6
- **Internationalization**: Custom i18n system
- **Build**: No build process required (development simplicity)

## Page Structure

### 1. Landing Page (index.html)

**Layout:**

- Hero section with compelling tagline
- Interactive map taking 70% of viewport
- Quick report button (floating action button)
- Language selector in header
- Statistics counter (total reports, resolved issues)

**Map Features:**

- Cluster markers at lower zoom levels
- Individual issue pins at higher zoom levels
- Click on pin to show issue preview popup
- Click "View Details" to open issue modal
- Current location detection and centering

**Quick Report Flow:**

1. Click floating "Report Issue" button
2. GPS location detection (with fallback to manual selection)
3. Category selection (visual icons)
4. Photo capture/upload with EXIF extraction
5. Optional description
6. Submit and get shareable link

### 2. Issue Details Modal

**Content:**

- Large image carousel
- Issue details (category, description, location)
- Status timeline with timestamps
- Share buttons (WhatsApp, Twitter, copy link)
- "Mark as resolved" button for public users
- Map showing exact location

### 3. Admin Dashboard (admin.html)

**Sections:**

- Summary statistics cards
- Reports management table with filters
- Image moderation queue
- User management (admin only)
- Bulk actions for reports

### 4. Moderation Panel

**Features:**

- Image approval/rejection interface
- Side-by-side comparison view
- Batch moderation capabilities
- Moderation history log

## UI Components

### Map Component

```javascript
// Map initialization with clustering
const map = L.map("map").setView([20.5937, 78.9629], 5); // India center

// Cluster configuration
const markers = L.markerClusterGroup({
  chunkedLoading: true,
  maxClusterRadius: 80,
});

// Custom cluster icon based on issue types
const createClusterIcon = (cluster) => {
  const count = cluster.getChildCount();
  const categories = getClusterCategories(cluster);
  return L.divIcon({
    html: `<div class="cluster-marker">${count}</div>`,
    className: `marker-cluster marker-cluster-${getMainCategory(categories)}`,
    iconSize: [40, 40],
  });
};
```

### Report Form Component

```javascript
const ReportForm = {
  init() {
    this.setupLocationDetection();
    this.setupPhotoCapture();
    this.setupCategorySelection();
  },

  setupLocationDetection() {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => this.setLocation(position),
        () => this.showLocationError()
      );
    }
  },

  extractEXIFLocation(file) {
    // Extract GPS coordinates from image EXIF data
    EXIF.getData(file, function () {
      const lat = EXIF.getTag(this, "GPSLatitude");
      const lon = EXIF.getTag(this, "GPSLongitude");
      if (lat && lon) {
        ReportForm.setLocation({ coords: { latitude: lat, longitude: lon } });
      }
    });
  },
};
```

## Styling Guidelines

### Color Scheme

```css
:root {
  --primary-color: #2563eb; /* Blue */
  --secondary-color: #10b981; /* Green */
  --danger-color: #ef4444; /* Red */
  --warning-color: #f59e0b; /* Amber */
  --background: #f8fafc; /* Light gray */
  --surface: #ffffff; /* White */
  --text-primary: #1f2937; /* Dark gray */
  --text-secondary: #6b7280; /* Medium gray */
}
```

### Category Icons and Colors

```css
.category-potholes {
  color: #8b5cf6;
}
.category-streetlight {
  color: #f59e0b;
}
.category-water {
  color: #3b82f6;
}
.category-drainage {
  color: #06b6d4;
}
.category-sidewalk {
  color: #6b7280;
}
.category-accident {
  color: #ef4444;
}
.category-garbage {
  color: #84cc16;
}
.category-traffic {
  color: #f97316;
}
```

### Responsive Design

```css
/* Mobile-first approach */
.container {
  padding: 1rem;
}

@media (min-width: 768px) {
  .container {
    padding: 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }
}

/* Map responsive height */
#map {
  height: 60vh;
}

@media (min-width: 768px) {
  #map {
    height: 70vh;
  }
}
```

## Internationalization

### Language Structure

```javascript
const translations = {
  en: {
    "report.issue": "Report Issue",
    "categories.potholes": "Potholes",
    "status.pending": "Pending",
    "upload.photo": "Upload Photo",
  },
  hi: {
    "report.issue": "समस्या की रिपोर्ट करें",
    "categories.potholes": "गड्ढे",
    "status.pending": "लंबित",
    "upload.photo": "फोटो अपलोड करें",
  },
};

// Usage
const t = (key) => translations[currentLanguage][key] || key;
```

### Language Detection

```javascript
// Detect browser language
const detectLanguage = () => {
  const browserLang = navigator.language.split("-")[0];
  const supportedLangs = ["en", "hi"];
  return supportedLangs.includes(browserLang) ? browserLang : "en";
};
```

## Interactive Features

### Real-time Updates

```javascript
// WebSocket connection for real-time report updates
const ws = new WebSocket("ws://localhost:8080/ws");

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  if (data.type === "new_report") {
    addReportToMap(data.report);
    updateStatistics();
  }
};
```

### Share Functionality

```javascript
const shareReport = (reportId) => {
  const url = `${window.location.origin}/reports/${reportId}`;

  if (navigator.share) {
    navigator.share({
      title: "Infrastructure Issue Report",
      text: "Check out this reported issue in our area",
      url: url,
    });
  } else {
    // Fallback to copy to clipboard
    navigator.clipboard.writeText(url);
    showNotification("Link copied to clipboard");
  }
};
```

### Photo Capture

```javascript
const startCamera = () => {
  navigator.mediaDevices
    .getUserMedia({
      video: { facingMode: "environment" }, // Rear camera
    })
    .then((stream) => {
      const video = document.getElementById("camera-preview");
      video.srcObject = stream;
    })
    .catch((err) => {
      console.error("Camera access denied:", err);
      showFileUpload(); // Fallback to file upload
    });
};
```

## Performance Optimizations

### Map Performance

- Use clustering for large datasets
- Lazy load issue details
- Implement viewport-based loading
- Cache map tiles locally

### Image Optimization

- Compress images before upload
- Progressive JPEG loading
- Lazy load images in issue list
- Thumbnail generation

### Bundle Size

- Use CDN for external libraries
- Minimize custom CSS/JS
- Compress static assets
- Implement service worker for caching

## Accessibility

- ARIA labels for interactive elements
- Keyboard navigation support
- High contrast mode support
- Screen reader compatibility
- Focus management in modals
