# Job Openings API

A RESTful API built with Go (Gin) for managing job openings. The API allows you to create, read, update, and delete job opening listings with features like role, company, location, salary, and remote work status.

## Features

- CRUD operations for job openings
- SQLite database with GORM
- Swagger documentation
- Input validation
- Structured logging
- RESTful API design

## Technologies

- Go 1.23.4
- Gin Web Framework
- GORM with SQLite
- Swagger (gin-swagger)
- Validator

## Installation

1. Clone the repository:

```
git clone git@github.com:MaiconGilton/go-job-openings.git
cd job-openings
```

2. Install dependencies:

```
go mod download
```

3. Run the application:

```
go run main.go
```

The server will start on `http://localhost:8080`

## API Documentation

Once the server is running, you can access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

### Available Endpoints

- `GET /api/v1/opening?id={id}` - Get a specific job opening
- `POST /api/v1/opening` - Create a new job opening
- `PUT /api/v1/opening` - Update an existing job opening
- `DELETE /api/v1/opening?id={id}` - Delete a job opening
- `GET /api/v1/openings` - List all job openings

## Project Structure

```
.
├── config/         # Configuration files (database, logger)
├── docs/          # Swagger documentation
├── handler/       # HTTP handlers
├── router/        # Route definitions
├── schemas/       # Data models
├── main.go        # Application entry point
└── README.md
```

## Example Request

Creating a new job opening:

```bash
curl -X POST http://localhost:8080/api/v1/opening \
  -H 'Content-Type: application/json' \
  -d '{
    "role": "Software Engineer",
    "company": "Tech Corp",
    "location": "New York",
    "salary": 120000,
    "link": "https://example.com/job",
    "remote": true
  }'
```

## Database

The application uses SQLite as its database. The database file will be automatically created at `./db/main.db` when you first run the application.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
