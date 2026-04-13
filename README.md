# Address Book Spring Boot API

This project is a **Spring Boot REST API backend** for storing contact addresses in a MySQL database.  
It is designed to be consumed by any front-end UI (web, mobile, desktop).

## Tech Stack

- Java 17
- Spring Boot 3
- Spring Web
- Spring Data JPA
- MySQL
- Maven

## Database Table

The API persists data in an `address` table with the following columns:

- `name` (primary key)
- `address_line1`
- `city`
- `state`
- `country`

## API Endpoints

Base URL: `/api/addresses`

- `GET /api/addresses` → list all addresses
- `GET /api/addresses/{name}` → get one address by name
- `POST /api/addresses` → create a new address
- `PUT /api/addresses/{name}` → update an address
- `DELETE /api/addresses/{name}` → delete an address

### Sample POST Request

```json
{
  "name": "John Doe",
  "addressLine1": "123 Main Street",
  "city": "Austin",
  "state": "Texas",
  "country": "USA"
}
```

## MySQL Configuration

Application properties use environment variable overrides:

- `DB_URL` (default: `jdbc:mysql://localhost:3306/addressbook`)
- `DB_USERNAME` (default: `root`)
- `DB_PASSWORD` (default: `root`)
- `SERVER_PORT` (default: `8080`)

## Run the Application

```bash
mvn spring-boot:run
```

## Run Tests

```bash
mvn test
```

Tests use an in-memory H2 database.
