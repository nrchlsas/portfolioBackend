# Portfolio Backend API

Backend RESTful API untuk aplikasi Portfolio yang dibangun dengan Go (Golang), MySQL, dan dokumentasi Swagger/OpenAPI.

## 🚀 Quick Start

### Prerequisites
- Go 1.21 atau lebih tinggi
- MySQL 8.0 atau lebih tinggi
- Docker & Docker Compose (optional, untuk MySQL setup yang mudah)

### Setup

1. **Clone dan masuk ke direktori backend:**
```bash
cd golang-backend
```

2. **Buat .env file:**
```bash
cp .env.example .env
```

3. **Update nilai .env sesuai konfigurasi Anda:**
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password_here
DB_NAME=portfolio_db
SERVER_PORT=8080
```

4. **Setup MySQL Database:**

**Option A: Menggunakan Docker Compose (Recommended)**
```bash
make db-up
```

**Option B: Manual Setup**
```bash
mysql -u root -p < database_schema.sql
```

5. **Install dependencies:**
```bash
make install
```

6. **Jalankan server:**
```bash
make run
```

Server akan berjalan di `http://localhost:8080`

## 📚 Available Endpoints

### Health Check
- `GET /api/health` - Check API status

### Skills
- `GET /api/skills` - Get semua skills
- `GET /api/skills/:id` - Get skill berdasarkan ID
- `POST /api/skills` - Tambah skill baru
- `PUT /api/skills/:id` - Update skill
- `DELETE /api/skills/:id` - Hapus skill

### Experience
- `GET /api/experience` - Get semua experience
- `GET /api/experience/:id` - Get experience berdasarkan ID
- `POST /api/experience` - Tambah experience baru
- `PUT /api/experience/:id` - Update experience
- `DELETE /api/experience/:id` - Hapus experience

### Services
- `GET /api/services` - Get semua services
- `GET /api/services/:id` - Get service berdasarkan ID
- `POST /api/services` - Tambah service baru
- `PUT /api/services/:id` - Update service
- `DELETE /api/services/:id` - Hapus service

### Projects
- `GET /api/projects` - Get semua projects
- `GET /api/projects/:id` - Get project berdasarkan ID
- `POST /api/projects` - Tambah project baru
- `PUT /api/projects/:id` - Update project
- `DELETE /api/projects/:id` - Hapus project

### Contact
- `GET /api/contact` - Get contact information
- `PUT /api/contact` - Update contact information

## 📖 API Documentation

Swagger UI tersedia di: `http://localhost:8080/swagger/index.html`

Untuk generate/update Swagger docs:
```bash
make swagger
```

## 📁 Project Structure

```
golang-backend/
├── main.go                 # Entry point aplikasi
├── config/
│   └── database.go        # Database configuration
├── models/
│   └── models.go          # Data models dan request/response structs
├── handlers/
│   ├── skill_handler.go
│   ├── experience_handler.go
│   ├── service_handler.go
│   ├── project_handler.go
│   └── contact_handler.go
├── routes/
│   └── routes.go          # Route definitions
├── docs/
│   └── docs.go            # Generated Swagger docs
├── .env.example           # Environment variables template
├── docker-compose.yml     # Docker setup untuk MySQL
├── database_schema.sql    # SQL schema dan sample data
├── Makefile              # Build commands
├── go.mod                # Go module dependencies
└── README.md             # Documentation
```

## 🛠️ Available Commands

```bash
# Setup dan install dependencies
make install

# Jalankan development server
make run

# Build binary
make build

# Run tests
make test

# Clean build artifacts
make clean

# Start MySQL dengan Docker
make db-up

# Stop MySQL
make db-down

# Generate Swagger docs
make swagger

# Run dengan hot reload (development)
make dev

# Setup semua (install, db, swagger, run)
make all
```

## 📝 Sample API Calls

### Health Check
```bash
curl http://localhost:8080/api/health
```

### Get All Skills
```bash
curl http://localhost:8080/api/skills
```

### Create a Skill
```bash
curl -X POST http://localhost:8080/api/skills \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Go Programming"
  }'
```

### Get All Projects
```bash
curl http://localhost:8080/api/projects
```

### Create Project
```bash
curl -X POST http://localhost:8080/api/projects \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My New Project",
    "description": "Project description",
    "tags": ["Go", "React"],
    "gradient": "from-blue-500 to-purple-500",
    "year": 2024,
    "emoji": "🚀",
    "link_color": "blue"
  }'
```

### Get Contact Info
```bash
curl http://localhost:8080/api/contact
```

### Update Contact Info
```bash
curl -X PUT http://localhost:8080/api/contact \
  -H "Content-Type: application/json" \
  -d '{
    "email": "nurcholisas123@gmail.com",
    "phone": "+62 859-5960-1389",
    "location": "Jakarta, Indonesia",
    "social": [
      {
        "name": "LinkedIn",
        "url": "https://linkedin.com/in/nurcholis-ahmad-syarif",
        "initials": "in",
        "gradient": "from-blue-600 to-blue-700"
      }
    ]
  }'
```

## 🔗 Frontend Integration

Backend ini dirancang untuk bekerja dengan Next.js portfolio frontend. Frontend dapat mengakses semua endpoint melalui:

```javascript
const API_BASE = 'http://localhost:8080/api'

// Example dalam React
const fetchSkills = async () => {
  const response = await fetch(`${API_BASE}/skills`)
  return response.json()
}
```

## 🐛 Troubleshooting

### Database Connection Error
- Pastikan MySQL server running
- Cek konfigurasi `.env` file
- Verifikasi credentials database

### Port Already in Use
Ubah `SERVER_PORT` di `.env` file

### Swagger Docs Not Loading
```bash
go install github.com/swaggo/swag/cmd/swag@latest
make swagger
```

### Missing Dependencies
```bash
go mod download
go mod tidy
```

## 📦 Dependencies

- `github.com/gin-gonic/gin` - Web framework
- `gorm.io/gorm` - ORM
- `gorm.io/driver/mysql` - MySQL driver
- `github.com/joho/godotenv` - Environment variables
- `github.com/swaggo/gin-swagger` - Swagger UI
- `github.com/swaggo/swag` - Swagger docs generator

## 📄 Database Schema

Lihat [database_schema.sql](./database_schema.sql) untuk schema lengkap dan sample data.

Atau baca [DB_SETUP.md](./DB_SETUP.md) untuk instruksi setup database detailed.

## 🚀 Deployment

Untuk development: `make run`

Untuk production build:
```bash
make build
./bin/portfolio-backend
```

## 📞 Support

Untuk pertanyaan atau issues, silakan buat issue di repository.

## 📄 License

MIT License
