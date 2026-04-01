# Portfolio Backend - Complete Setup Guide

## 📋 Table of Contents
1. [Prerequisites](#prerequisites)
2. [Quick Start (5 minutes)](#quick-start-5-minutes)
3. [Detailed Setup](#detailed-setup)
4. [Running the Backend](#running-the-backend)
5. [Database Management](#database-management)
6. [API Documentation](#api-documentation)
7. [Troubleshooting](#troubleshooting)
8. [Frontend Integration](#frontend-integration)

---

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21 or higher**
  ```bash
  go version  # Check your version
  ```
  [Download Go](https://golang.org/dl/)

- **MySQL 8.0 or higher**
  ```bash
  mysql --version  # Check your version
  ```
  [Download MySQL](https://dev.mysql.com/downloads/mysql/)

- **Docker (Optional but recommended for easier MySQL setup)**
  ```bash
  docker --version
  docker-compose --version
  ```
  [Download Docker](https://www.docker.com/)

---

## Quick Start (5 minutes)

### 1. Setup Environment
```bash
cd golang-backend
cp .env.example .env
```

### 2. Update .env file
Edit `.env` and update database credentials:
```bash
# For local/Docker MySQL, use these defaults
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password_here
DB_NAME=portfolio_db
SERVER_PORT=8080
```

### 3. Start MySQL (using Docker)
```bash
make db-up
```

Wait for MySQL to be ready (about 10 seconds)

### 4. Install Dependencies
```bash
make install
```

### 5. Run the Backend
```bash
make run
```

Server will start at `http://localhost:8080`

### 6. Access Swagger Documentation
Open in browser: `http://localhost:8080/swagger/index.html`

---

## Detailed Setup

### Step 1: Clone/Navigate to Backend Folder
```bash
cd /Users/macbookpro/Documents/SourceCode/Portopolio/golang-backend
```

### Step 2: Environment Configuration

Create `.env` file from template:
```bash
cp .env.example .env
```

Edit `.env` with your database credentials:
```bash
# Database Configuration
DB_HOST=localhost          # MySQL host
DB_PORT=3306              # MySQL port
DB_USER=root              # MySQL username
DB_PASSWORD=root          # MySQL password
DB_NAME=portfolio_db      # Database name
SERVER_PORT=8080          # Backend server port
```

### Step 3: Database Setup

**Option A: Using Docker Compose (Recommended)**
```bash
# Start MySQL container with database
make db-up

# Verify MySQL is running
docker ps | grep portfolio_mysql

# Stop MySQL container
make db-down
```

**Option B: Manual MySQL Setup**
```bash
# Connect to MySQL
mysql -u root -p

# In MySQL prompt, run:
source database_schema.sql;

# Exit MySQL
EXIT;
```

**Option C: From MySQL Command Line**
```bash
mysql -u root -p < database_schema.sql
```

### Step 4: Install Dependencies

```bash
# Download and verify all Go dependencies
make install

# Or manually:
go mod download
go mod tidy
go install github.com/swaggo/swag/cmd/swag@latest
```

### Step 5: Verify Setup

Check database connection:
```bash
# Connect to MySQL
mysql -u root -p -e "USE portfolio_db; SHOW TABLES;"
```

You should see these tables:
- contact
- experience
- projects
- services
- skills

---

## Running the Backend

### Development Mode
```bash
# Simple run
make run

# Or manually
go run main.go
```

### Build Binary
```bash
make build

# Run binary
./bin/portfolio-backend
```

### With Hot Reload (Optional)
Install air for hot reload:
```bash
go install github.com/cosmtrek/air@latest

# Then run with air watching for changes
air
```

### Expected Output
```
Loading .env
Server running on http://localhost:8080
Swagger docs available at http://localhost:8080/swagger/index.html
```

---

## Database Management

### Backup Database
```bash
mysqldump -u root -p portfolio_db > backup_portfolio.sql
```

### Restore Database
```bash
mysql -u root -p portfolio_db < backup_portfolio.sql
```

### Reset Database (Delete all data)
```bash
mysql -u root -p -e "DROP DATABASE IF EXISTS portfolio_db; CREATE DATABASE portfolio_db;"
mysql -u root -p < database_schema.sql
```

### View Database Tables
```bash
mysql -u root -p portfolio_db -e "SHOW TABLES;"
```

### View Table Structure
```bash
mysql -u root -p portfolio_db -e "DESCRIBE skills;"
```

---

## API Documentation

### Swagger UI
Access at: `http://localhost:8080/swagger/index.html`

### Generate/Update Swagger Docs
```bash
make swagger

# Manually:
swag init
```

### API Documentation Files
- [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) - Complete endpoint reference
- [DB_SETUP.md](./DB_SETUP.md) - Database setup details

### Common Endpoints

#### Health Check
```bash
curl http://localhost:8080/api/health
```

#### Get All Skills
```bash
curl http://localhost:8080/api/skills
```

#### Create Skill
```bash
curl -X POST http://localhost:8080/api/skills \
  -H "Content-Type: application/json" \
  -d '{"name": "New Skill"}'
```

#### Get Contact
```bash
curl http://localhost:8080/api/contact
```

See [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) for more examples.

---

## Troubleshooting

### Issue: Connection Refused
**Problem:** `dial tcp [::1]:3306: connect: connection refused`

**Solutions:**
1. Check MySQL is running:
   ```bash
   # Docker
   docker ps | grep portfolio_mysql
   
   # Manual MySQL
   mysql -u root -p -e "SELECT 1;"
   ```

2. Restart MySQL:
   ```bash
   # Docker
   make db-down
   make db-up
   
   # Manual (macOS)
   brew services restart mysql
   ```

3. Check credentials in `.env` file

### Issue: Port Already in Use
**Problem:** `listen tcp :8080: bind: address already in use`

**Solutions:**
1. Change port in `.env`:
   ```
   SERVER_PORT=8081
   ```

2. Kill process on port 8080:
   ```bash
   # Find process
   lsof -i :8080
   
   # Kill it
   kill -9 <PID>
   ```

### Issue: Database Not Found
**Problem:** `Error 1049: Unknown database 'portfolio_db'`

**Solutions:**
1. Verify database was created:
   ```bash
   mysql -u root -p -e "SHOW DATABASES;"
   ```

2. Create database manually:
   ```bash
   mysql -u root -p < database_schema.sql
   ```

### Issue: Go Modules Error
**Problem:** `go: missing go.sum entry...`

**Solutions:**
```bash
go mod tidy
go mod download
```

### Issue: Dependencies Not Installed
**Problem:** `cannot find package...`

**Solutions:**
```bash
go mod install
# Or
make install
```

### Issue: Swagger Docs Not Loading
**Problem:** Swagger UI returns 404

**Solutions:**
1. Generate Swagger docs:
   ```bash
   make swagger
   ```

2. Make sure swag is installed:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

### Issue: Build Fails
**Problem:** Compilation errors

**Solutions:**
```bash
# Clean and rebuild
make clean
make install
make build
```

---

## Frontend Integration

### Configure Frontend

Update your Next.js frontend `.env.local`:
```
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

### Example React Hook
```javascript
// hooks/usePortfolioAPI.ts
import { useEffect, useState } from 'react'

export function usePortfolioAPI() {
  const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'

  const [skills, setSkills] = useState([])
  const [projects, setProjects] = useState([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function fetchData() {
      try {
        const [skillsRes, projectsRes] = await Promise.all([
          fetch(`${API_BASE}/skills`),
          fetch(`${API_BASE}/projects`)
        ])

        const skillsData = await skillsRes.json()
        const projectsData = await projectsRes.json()

        setSkills(skillsData.data || [])
        setProjects(projectsData.data || [])
      } finally {
        setLoading(false)
      }
    }

    fetchData()
  }, [])

  return { skills, projects, loading }
}
```

### Update Frontend Components
```javascript
// components/sections/SkillsSection.tsx
import { usePortfolioAPI } from '@/hooks/usePortfolioAPI'

export function SkillsSection() {
  const { skills } = usePortfolioAPI()

  return (
    <section>
      {skills.map(skill => (
        <div key={skill.id}>{skill.name}</div>
      ))}
    </section>
  )
}
```

---

## Available Commands

```bash
# Show all available commands
make help

# Install dependencies
make install

# Run development server
make run

# Build binary
make build

# Run tests
make test

# Clean build artifacts
make clean

# Database operations
make db-up      # Start MySQL
make db-down    # Stop MySQL

# Generate Swagger docs
make swagger

# Full setup
make all        # install + db-up + swagger + run
```

---

## Project Structure

```
golang-backend/
├── main.go                  # Application entry point
├── config/
│   └── database.go         # Database connection setup
├── models/
│   └── models.go           # Data models and DTOs
├── handlers/
│   ├── skill_handler.go
│   ├── experience_handler.go
│   ├── service_handler.go
│   ├── project_handler.go
│   └── contact_handler.go  # Business logic
├── routes/
│   └── routes.go           # Route definitions
├── docs/
│   └── docs.go             # Generated Swagger docs
├── .env.example            # Environment template
├── .env                    # Environment variables (git ignored)
├── docker-compose.yml      # MySQL Docker setup
├── database_schema.sql     # Database schema + sample data
├── Makefile               # Build commands
├── go.mod                 # Module dependencies
├── go.sum                 # Dependency checksums
├── README.md              # Quick reference
├── DB_SETUP.md            # Database setup guide
├── API_DOCUMENTATION.md   # API reference
└── SETUP_GUIDE.md         # This file
```

---

## Next Steps

1. **Start the backend:**
   ```bash
   make run
   ```

2. **Access Swagger UI:**
   Open http://localhost:8080/swagger/index.html

3. **Test API endpoints:**
   See [API_DOCUMENTATION.md](./API_DOCUMENTATION.md)

4. **Integrate with frontend:**
   Update Next.js frontend to fetch data from this backend

5. **Deploy:**
   Build binary with `make build` and deploy to your server

---

## Common Tasks

### Add New Entity (e.g., Testimonials)

1. **Add model in `models/models.go`**
```go
type Testimonial struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Author    string    `json:"author"`
    Message   string    `json:"message"`
    Rating    int       `json:"rating"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

2. **Create handler `handlers/testimonial_handler.go`**

3. **Add routes in `routes/routes.go`**

4. **Update database schema in `database_schema.sql`**

5. **Run migration:**
```go
db.AutoMigrate(&models.Testimonial{})
```

### Change Database Credentials

Edit `.env` file:
```
DB_USER=newuser
DB_PASSWORD=newpassword
```

Then restart the application.

### Deploy to Production

1. Build:
   ```bash
   make build
   ```

2. Set environment variables on server:
   ```bash
   export DB_HOST=production-db-host
   export DB_USER=prod_user
   export DB_PASSWORD=strong_password
   export DB_NAME=portfolio_db_prod
   export SERVER_PORT=8080
   ```

3. Run binary:
   ```bash
   ./bin/portfolio-backend
   ```

---

## Support & Documentation

- Go Documentation: https://golang.org/doc/
- Gin Framework: https://github.com/gin-gonic/gin
- GORM: https://gorm.io/
- MySQL: https://dev.mysql.com/doc/
- Swagger: https://swagger.io/

---

## License

MIT License - See LICENSE file for details
