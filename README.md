# Portfolio Backend API

This is a backend API for a portfolio website built with Go Fiber and Supabase, using GORM for database operations.

## Setup

1. Clone the repository
2. Create a `.env` file in the root directory with the following variables:
   ```
   SUPABASE_URL=your_supabase_project_url
   SUPABASE_DB_PASSWORD=your_supabase_database_password
   PORT=3000
   ```

   Note: The `SUPABASE_DB_PASSWORD` is your database password from Supabase dashboard (Database Settings > Database Password), not the anon key.

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

   The application will automatically create the following tables in your Supabase database:
   - `experiences`
   - `projects`
   - `skill_categories`

## Deployment

1. Build the application:
   ```bash
   go build -o wannn-site-rebuild-api main.go
   ```
2. Run the application:
   ```bash
   ./wannn-site-rebuild-api

3. Create a systemd service file:
   ```bash
   sudo nano /etc/systemd/system/wannn-site-rebuild-api.service
   ```
4. Add the following content to the service file:
   ```bash
    [Unit]
    Description=Wannn Site API Service
    After=network.target

    [Service]
    Type=simple
    User=ikhw
    WorkingDirectory=/home/ikhw/wannn-site-rebuild-api
    ExecStart=/home/ikhw/wannn-site-rebuild-api/wannn-site-rebuild-api
    Restart=always
    RestartSec=3
    Environment=GO_ENV=production
    Environment=PORT=3000

    [Install]
    WantedBy=multi-user.target
    ```

5. Reload systemd:
   ```bash
   sudo systemctl daemon-reload
   ```
6. Enable and start the service:
   ```bash
   sudo systemctl enable wannn-site-rebuild-api
## API Endpoints

### Experiences
- GET `/api/experiences` - Get all experiences
- GET `/api/experiences/:id` - Get experience by ID
- POST `/api/experiences` - Create new experience
- PUT `/api/experiences/:id` - Update experience
- DELETE `/api/experiences/:id` - Delete experience

Example Experience JSON:
```json
{
  "title": "Backend Developer Intern",
  "company": "Tech Company",
  "period": "Jun 2023 - Dec 2023",
  "description": [
    "Developed and maintained RESTful APIs",
    "Implemented database optimizations",
    "Collaborated with frontend team"
  ]
}
```

### Projects
- GET `/api/projects` - Get all projects
- GET `/api/projects/:id` - Get project by ID
- POST `/api/projects` - Create new project
- PUT `/api/projects/:id` - Update project
- DELETE `/api/projects/:id` - Delete project

Example Project JSON:
```json
{
  "title": "E-Commerce Backend",
  "description": "A scalable backend system for an e-commerce platform",
  "technologies": ["Node.js", "Express", "PostgreSQL", "Redis"],
  "link": "https://github.com/username/project"
}
```

### Skill Categories
- GET `/api/skills` - Get all skill categories
- GET `/api/skills/:id` - Get skill category by ID
- POST `/api/skills` - Create new skill category
- PUT `/api/skills/:id` - Update skill category
- DELETE `/api/skills/:id` - Delete skill category

Example Skill Category JSON:
```json
{
  "title": "Backend Development",
  "skills": ["Node.js", "Express", "NestJS", "RESTful APIs", "GraphQL"]
}
```

## Database Schema

The following tables will be automatically created:

### experiences
- ID (uint, primary key)
- CreatedAt (timestamp)
- UpdatedAt (timestamp)
- DeletedAt (timestamp, nullable)
- Title (varchar(255))
- Company (varchar(255))
- Period (varchar(100))
- Description (text[])

### projects
- ID (uint, primary key)
- CreatedAt (timestamp)
- UpdatedAt (timestamp)
- DeletedAt (timestamp, nullable)
- Title (varchar(255))
- Description (text)
- Technologies (text[])
- Link (varchar(255))

### skill_categories
- ID (uint, primary key)
- CreatedAt (timestamp)
- UpdatedAt (timestamp)
- DeletedAt (timestamp, nullable)
- Title (varchar(255))
- Skills (text[])

## Technologies Used

- Go Fiber
- GORM
- Supabase (PostgreSQL)
- godotenv 