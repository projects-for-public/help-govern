# API Documentation

## Base URL

`http://localhost:8080/api/v1`

## Authentication

- Most endpoints are public for anonymous reporting
- Admin/Moderator endpoints require JWT token in Authorization header
- Format: `Authorization: Bearer <token>`

## Public Endpoints

### GET /reports

Get reports with clustering and filtering.

**Query Parameters:**

- `lat` (float): Center latitude
- `lng` (float): Center longitude
- `zoom` (int): Map zoom level (affects clustering)
- `category` (string): Filter by category
- `status` (string): Filter by status
- `limit` (int): Max results (default: 100)

**Response:**

```json
{
  "reports": [
    {
      "id": 1,
      "category": "potholes",
      "latitude": 26.9124,
      "longitude": 75.7873,
      "description": "Large pothole on main road",
      "status": "pending",
      "created_at": "2025-06-29T10:00:00Z",
      "images": [
        {
          "id": 1,
          "url": "https://res.cloudinary.com/...",
          "moderation_status": "approved"
        }
      ]
    }
  ],
  "clusters": [
    {
      "latitude": 26.9124,
      "longitude": 75.7873,
      "count": 5,
      "zoom_level": 10
    }
  ]
}
```

### GET /reports/:id

Get specific report details with timeline.

**Response:**

```json
{
  "id": 1,
  "category": "potholes",
  "latitude": 26.9124,
  "longitude": 75.7873,
  "description": "Large pothole on main road",
  "status": "in_progress",
  "created_at": "2025-06-29T10:00:00Z",
  "images": [
    {
      "id": 1,
      "url": "https://res.cloudinary.com/...",
      "type": "report",
      "moderation_status": "approved"
    }
  ],
  "timeline": [
    {
      "status": "pending",
      "notes": "Report submitted",
      "updated_at": "2025-06-29T10:00:00Z"
    },
    {
      "status": "verified",
      "notes": "Issue verified by moderator",
      "updated_at": "2025-06-29T11:30:00Z"
    }
  ]
}
```

### POST /reports

Submit new report (anonymous).

**Request Body:**

```json
{
  "category": "potholes",
  "latitude": 26.9124,
  "longitude": 75.7873,
  "description": "Large pothole causing vehicle damage",
  "images": ["base64_encoded_image_1", "base64_encoded_image_2"]
}
```

**Response:**

```json
{
  "id": 123,
  "share_url": "/reports/123",
  "message": "Report submitted successfully"
}
```

### GET /categories

Get all active categories.

**Response:**

```json
{
  "categories": [
    {
      "id": 1,
      "name": "potholes",
      "name_hi": "गड्ढे",
      "description": "Road potholes and surface damage",
      "icon_class": "fas fa-road"
    }
  ]
}
```

## Authentication Endpoints

### POST /auth/login

Login for moderators/admins.

**Request Body:**

```json
{
  "username": "moderator1",
  "password": "password123"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "moderator1",
    "role": "moderator"
  }
}
```

## Protected Endpoints (Moderator/Admin)

### GET /admin/reports

Get reports for moderation dashboard.

**Query Parameters:**

- `status` (string): Filter by status
- `moderation_status` (string): Filter by image moderation status
- `page` (int): Page number
- `limit` (int): Results per page

### PUT /admin/reports/:id/status

Update report status.

**Request Body:**

```json
{
  "status": "verified",
  "notes": "Issue confirmed on-site"
}
```

### PUT /admin/images/:id/moderate

Moderate image approval.

**Request Body:**

```json
{
  "moderation_status": "approved",
  "notes": "Image is relevant and appropriate"
}
```

### POST /admin/users (Admin only)

Create new moderator account.

**Request Body:**

```json
{
  "username": "newmoderator",
  "email": "mod@example.com",
  "password": "temppassword",
  "role": "moderator"
}
```

## Rate Limiting

- Anonymous endpoints: 10 requests per minute per IP
- Authenticated endpoints: 100 requests per minute per user
- Image upload: 5 uploads per hour per IP

## Error Responses

```json
{
  "error": "Error message",
  "code": "ERROR_CODE",
  "details": "Additional error details"
}
```

**Common Error Codes:**

- `VALIDATION_ERROR` (400)
- `UNAUTHORIZED` (401)
- `FORBIDDEN` (403)
- `NOT_FOUND` (404)
- `RATE_LIMITED` (429)
- `INTERNAL_ERROR` (500)

## Image Upload Notes

- Images are uploaded as base64 strings in request body
- Automatic EXIF GPS extraction if available
- Images go through moderation before display
- Supported formats: JPEG, PNG, WebP
- Max size: 5MB per image
- Max 3 images per report
