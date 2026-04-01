# DB Setup Instructions

## Prerequisites
- MySQL Server 8.0 or higher installed and running
- MySQL Client tools installed

## Database Setup Steps

### 1. Login to MySQL
```bash
mysql -u root -p
```

### 2. Run the SQL Schema
```bash
mysql -u root -p < database_schema.sql
```

Or manually copy and paste the SQL from `database_schema.sql` into MySQL client.

### 3. Verify Database Creation
```sql
USE portfolio_db;
SHOW TABLES;
```

You should see:
- contact
- experience
- projects
- services
- skills

### 4. Verify Sample Data
```sql
SELECT * FROM skills LIMIT 5;
SELECT * FROM experience;
SELECT * FROM services;
SELECT * FROM projects;
SELECT * FROM contact;
```

## Environment Configuration

### 1. Create .env file
Copy `.env.example` to `.env`:
```bash
cp .env.example .env
```

### 2. Update .env with your database credentials
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password_here
DB_NAME=portfolio_db
SERVER_PORT=8080
```

## Running the Backend

### 1. Install Go Dependencies
```bash
go mod download
go mod tidy
```

### 2. Run the Server
```bash
go run main.go
```

### 3. Access the API
- Health Check: http://localhost:8080/api/health
- Swagger Documentation: http://localhost:8080/swagger/index.html

## API Endpoints Overview

### Skills
- `GET /api/skills` - Get all skills
- `POST /api/skills` - Create skill
- `GET /api/skills/:id` - Get skill by ID
- `PUT /api/skills/:id` - Update skill
- `DELETE /api/skills/:id` - Delete skill

### Experience
- `GET /api/experience` - Get all experience
- `POST /api/experience` - Create experience
- `GET /api/experience/:id` - Get experience by ID
- `PUT /api/experience/:id` - Update experience
- `DELETE /api/experience/:id` - Delete experience

### Services
- `GET /api/services` - Get all services
- `POST /api/services` - Create service
- `GET /api/services/:id` - Get service by ID
- `PUT /api/services/:id` - Update service
- `DELETE /api/services/:id` - Delete service

### Projects
- `GET /api/projects` - Get all projects
- `POST /api/projects` - Create project
- `GET /api/projects/:id` - Get project by ID
- `PUT /api/projects/:id` - Update project
- `DELETE /api/projects/:id` - Delete project

### Contact
- `GET /api/contact` - Get contact information
- `PUT /api/contact` - Update contact information

### Health
- `GET /api/health` - API health check

## Testing with cURL

### Test Health Check
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
  -d '{"name": "Go Programming"}'
```

### Get Contact Information
```bash
curl http://localhost:8080/api/contact
```

### Update Contact Information
```bash
curl -X PUT http://localhost:8080/api/contact \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "phone": "+62 859-5960-1389",
    "location": "Jakarta, Indonesia",
    "social": [
      {"name": "LinkedIn", "url": "https://linkedin.com/in/nurcholis-ahmad-syarif"}
    ]
  }'
```

## Troubleshooting

### Database Connection Error
Ensure MySQL is running and credentials in .env are correct:
- Check DB_HOST, DB_PORT, DB_USER, DB_PASSWORD

### Port Already in Use
If port 8080 is already in use, change SERVER_PORT in .env file

### Missing Dependencies
Run `go mod tidy` to ensure all dependencies are installed

### Swagger Docs Not Showing
Make sure you have swag CLI installed:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Then regenerate docs:
```bash
swag init
```
