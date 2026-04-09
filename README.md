# URL Shortener

A full-stack URL shortening service with authentication and click analytics.

## Tech Stack

| Layer | Technology |
|-------|------------|
| **Backend** | Go 1.25 |
| **Frontend** | Next.js 16, React 19 |
| **Database** | PostgreSQL 15 |
| **Cache** | Redis 7 |
| **Auth** | JWT |
| **Deployment** | Railway (Backend), Vercel (Frontend) |

## Features

- Shorten long URLs with auto-generated codes
- User authentication (register/login)
- Personal URL management (create, view, delete)
- Click tracking and analytics
- Redis caching for fast redirects
- Responsive UI with Tailwind CSS

## Project Structure

```
url-shortener/
├── cmd/api/main.go           # Backend entry point
├── internal/
│   ├── auth/                 # Authentication (JWT, password hashing)
│   ├── url/                  # URL CRUD operations
│   ├── redirect/             # Short URL redirect handler
│   └── analytics/            # Click tracking
├── pkg/
│   ├── database/             # PostgreSQL connection
│   └── cache/                # Redis cache
├── migrations/               # Database migrations
├── frontend/
│   ├── app/
│   │   ├── page.tsx          # Landing page
│   │   ├── login/page.tsx    # Login page
│   │   ├── register/page.tsx # Register page
│   │   └── dashboard/page.tsx # User dashboard
│   └── lib/api.ts            # API client
└── docker-compose.yaml       # Local development
```

## Prerequisites

- Docker & Docker Compose
- Go 1.25+
- Node.js 18+
- PostgreSQL 15 (or use Docker)
- Redis 7 (or use Docker)

## Local Development

### 1. Clone & Setup

```bash
git clone <repository-url>
cd url-shortener
```

### 2. Start Database & Cache

```bash
docker-compose up -d
```

### 3. Setup Backend

```bash
# Create .env file
cp .env.example .env

# Install dependencies
go mod tidy

# Run server
go run cmd/api/main.go
```

### 4. Setup Frontend

```bash
cd frontend

# Install dependencies
npm install

# Create .env.local
echo "NEXT_PUBLIC_API_URL=http://localhost:8080" > .env.local

# Run dev server
npm run dev
```

### 5. Access

- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

## Environment Variables

### Backend (.env)

| Variable | Description | Example |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `BASE_URL` | Production domain | `https://your-domain.com` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | `password` |
| `DB_NAME` | Database name | `urlshortener` |
| `REDIS_HOST` | Redis host | `localhost` |
| `REDIS_PORT` | Redis port | `6379` |
| `REDIS_PASSWORD` | Redis password | `password` |
| `JWT_SECRET` | JWT signing secret | `your-secret-key` |

### Frontend (.env.local)

| Variable | Description | Example |
|----------|-------------|---------|
| `NEXT_PUBLIC_API_URL` | Backend API URL | `https://api.your-domain.com` |

## API Endpoints

### Authentication

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/auth/register` | Register new user |
| POST | `/auth/login` | Login user |

### URL Management

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/urls` | Required | Create short URL |
| GET | `/my-urls` | Required | Get user's URLs |
| DELETE | `/urls/:id` | Required | Delete URL |
| GET | `/:shortCode` | No | Redirect to original URL |

### Analytics

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/analytics/:shortCode` | Required | Get click stats |

### Request/Response Examples

**Register:**
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

**Create Short URL:**
```bash
curl -X POST http://localhost:8080/urls \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"long_url":"https://example.com/very-long-url"}'
```

**Response:**
```json
{
  "id": 1,
  "short_code": "abc123",
  "short_url": "https://your-domain.com/abc123",
  "long_url": "https://example.com/very-long-url"
}
```

## Deployment

### Backend (Railway)

1. Connect GitHub repository to Railway
2. Add environment variables in Railway dashboard:
   - `BASE_URL`: Your Railway domain (e.g., `https://your-app.up.railway.app`)
   - `DB_*`: Railway PostgreSQL connection string
   - `REDIS_*`: Railway Redis connection string
   - `JWT_SECRET`: Generate secure random string
3. Railway auto-deploys on push to main branch

### Frontend (Vercel)

1. Import GitHub repository to Vercel
2. Add environment variable:
   - `NEXT_PUBLIC_API_URL`: Your Railway backend URL
3. Vercel auto-deploys on push to main branch

## Database Schema

```sql
-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- URLs table
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    long_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Clicks table (analytics)
CREATE TABLE clicks (
    id SERIAL PRIMARY KEY,
    url_id INTEGER REFERENCES urls(id),
    clicked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## License

[MIT License](LICENSE)
