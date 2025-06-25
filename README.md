# Go Ticketing Server

A backend microservice built in Go to manage employee support tickets, designed to integrate seamlessly with a RAG-based chatbot for enhanced employee assistance.

## Features

* **Raise Tickets:** Allows employees to log new support issues with a detailed description and title, automatically associating them with a user (creating the user if they don't exist).
* **Get Tickets:** Retrieve a list of tickets for a specific user.

## Technologies Used

* **GoLang:** Backend language
* **Gin Gonic:** Web framework for building REST APIs
* **GORM:** ORM for database interactions
* **PostgreSQL:** Relational database for data persistence
* **`google/uuid`:** For UUID generation and handling
* **`go-playground/validator`:** For request validation

## Setup & Run

### Prerequisites

* Go (1.18+)
* PostgreSQL (running locally or accessible)

### Database Configuration

Update the following constants in `database/database.go` with your PostgreSQL credentials:

```go
const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "your_postgres_password" // Change this
  dbname   = "mockTicket"
)
```

Ensure your `mockTicket` database exists in PostgreSQL. The server will handle schema migration automatically via GORM's `AutoMigrate`.

### Run the Server

1. **Clone the repository** (if applicable, or navigate to your project root).
2. **Download dependencies:** `go mod tidy`
3. **Run the application:** `go run main.go`

The server will start on port `8080` by default (or the `PORT` environment variable if set).

## API Endpoints

### 1. Raise a Ticket (POST)

Creates a new support ticket.

* **URL:** `/raise-ticket`
* **Method:** `POST`
* **Headers:** `Content-Type: application/json`
* **Request Body Example:**

```json
{
  "user_id": "c1a2b3d4-e5f6-7890-1234-567890abcdef",
  "title": "Urgent: Laptop Display Issues and Software Glitches",
  "issue": "My laptop display has started flickering intermittently, especially when moving the screen.
            Additionally, several applications, including our internal CRM software, are crashing unexpectedly,
            causing significant disruption to my daily workflow and productivity. I have tried restarting the
            system multiple times, but the problem persists."
}
```
* **Success Response (201 Created):**

```json
{
  "ticket_id": 1,
  "user_id": "c1a2b3d4-e5f6-7890-1234-567890abcdef",
  "message": "Ticket raised successfully.",
  "status": "Raised",
  "created_at": "2025-06-26T10:00:00Z"
}
```

### 2. Get User's Tickets (GET)

Retrieves tickets for a specific user.

* **URL:** `/users/:user_id/tickets`

* **Method:** `GET`

* **Path Parameters:**

* `:user_id` (string, UUID): The ID of the user whose tickets are to be retrieved.



