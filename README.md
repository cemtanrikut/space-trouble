# SpaceTrouble

SpaceTrouble is an API that allows users to book tickets to various destinations within our solar system. In this project, SpaceTrouble shares launchpads with its competitor, SpaceX. Only one launch per day can take place from each launchpad. The API checks for conflicts with SpaceX launches and allows bookings only when there is no conflict.

## Features

- Book trips to specific destinations
- List all existing bookings
- Delete bookings
- Check SpaceX launch schedule to avoid conflicts

## Requirements

- **Golang** (v1.16 or higher)
- **PostgreSQL** (v13 or higher)
- **Docker and Docker Compose** (optional for containerized deployment)

## Setup

1. **Clone the Project**

   ```bash
   git clone https://github.com/cemtanrikut/spacetrouble.git
   cd spacetrouble
   ```

2. **Set Environment Variables**

Add the following environment variables to a .env file or export them in your terminal:

```bash
DB_CONN_STRING=postgresql://username:password@localhost:5432/spacetrouble?sslmode=disable
```

3. **Run with Docker (Recommended)**

To run the project in a containerized environment, use Docker Compose:

```bash
docker-compose up --build
```

4. **Run Database Migrations**

Run the necessary migration files to create database schemas. (Customize this step based on your migration files.

## Usage

### API Endpoints

| Endpoint                        | Method | Description                          |
|---------------------------------|--------|--------------------------------------|
| `/bookings`                     | POST   | Create a new booking.                |
| `/bookings`                     | GET    | List all bookings.                   |
| `/bookings/{id}`                | DELETE | Delete a specific booking.           |
| `/api/spacex/launchpads`        | GET    | List available launchpads.           |
| `/api/spacex/launches/upcoming` | GET    | List upcoming launches.              |

### Example Requests

1. Create a Booking

```http
POST /bookings
Content-Type: application/json

{
    "firstName": "John",
    "lastName": "Doe",
    "gender": "Male",
    "birthday": "1990-01-01",
    "launchpadID": "1",
    "destinationID": "Mars",
    "launchDate": "2049-12-31"
}
```

2. List All Bookings

```http
GET /bookings
```

3. Delete a Booking

```http
DELETE /bookings/{id}
```

## Project Structure

- /cmd: Main application entry point
- /internal/app: API handlers and routing
- /internal/service: Business logic and service layer
- /internal/repository: Database operations
- /internal/spacex: SpaceX API integration
- /pkg/db: Database connection configuration

## Tests

To run unit and functional tests:

```bash
go test ./...
```

## Contributing

If you would like to contribute, please open an issue first or submit a pull request.

## License

This project is licensed under the MIT License.