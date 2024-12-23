# Project Overview
This project is a RESTful API implemented in Go, designed to support user management, feed creation, and feed-follow functionality using PostgreSQL as the database. The project is structured into multiple phases to implement core features incrementally.

---

## HTTP Server Initialization

- **Initialized HTTP Server** using the `chi` package for routing.
- **Health Check Endpoint**: A simple health check API to verify that the server is running.

---

## Database Configuration

- **Configured PostgreSQL**:
  - Created a local PostgreSQL database.
  - Setup database connection URL: 
    ```
    postgres://<your_username>:<your_password>@localhost:5432/<your_db_name>?sslmode=disable
    ```
- **Configured `sqlc.yaml`**:
  - Set up the `queries` and `schema` directories within the `sql` folder.
- **Database Connection**:
  - Established connection with PostgreSQL.
  - Configured `connectionMaxIdleTime`.
  - Passed the connection pointer to `handlers/handleDBconn.go`.

---

## Create APIs

### For each API handler, Database Setup
- **Schema Migration**:
  - Created file i.e. `001_user.sql` in `sql/schema` for creating, dropping and migrating of DB tables the i.e. `user` table. This creates / drops table in database (here postgres)
  - Ran migrations using Goose:
    ```
    cd sql/schema
    goose postgres CONN up
    ```
    *`CONN` is the database connection URL excluding the `sslmode` flag.*

- **Query Generation**:
  - Created sql query file i.e. `user.sql` in i.e. `sql/queries`. This generated go compiled query fxns file i.e. `users.sql.go` in lets say `internal/database` which will be used to call queries fxn from go code.
  - Generated Go code using:
    ```
    sqlc generate
    ```

### User APIs Endpoints 
  - `POST /createUser`: Create a user.
  - `GET /getUserByApiKey`: Retrieve user by API key.

### Feed APIs Endpoints
  - `POST /createFeed`: Create a feed for a user.

### Feed-Follows APIs Endpoints
  - `POST /createFeedFollow`: Creating a feed_follow indicates that a user is now following a feed.
  - `GET /getAllFeedFollow`: Retrieving all feed followed by a user.
  - `DELETE /deleteFeedFollow/{feed_follow_id}`: Deleting a feed_follow is the same as “unfollowing” a feed.

---

## Getting Started

### Prerequisites
- Go 1.20+
- PostgreSQL 14+
- Goose CLI
- sqlc CLI

### Installation
1. Clone the repository:
    ```
    git clone <repository-url>
    cd <repository-folder>
    ```

2. Install dependencies:
    ```
    go mod tidy
    ```

3. Configure PostgreSQL and update the connection URL in the `.env` file.

4. Run migrations:
    ```
    cd sql/schema
    goose postgres CONN up
    ```

5. Generate query code:
    ```
    sqlc generate
    ```

### Running the Server
Start the HTTP server:
```bash
go run main.go
```

