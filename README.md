#  Todo API (Go + SQLite)

A simple RESTful Todo API built with **Go**, using **Chi router** and **SQLite** for persistence.

This project was created to learn backend fundamentals like:

* REST API design
* CRUD operations
* layered architecture (handlers → services → database)
* working with SQL in Go

---

##  Features

* ✅ Create, read, update, delete todos (CRUD)
* ✅ SQLite database (persistent storage)
* ✅ Clean project structure (handlers, services, db)
* ✅ JSON API with proper HTTP status codes
* ✅ Error handling with consistent responses

---

##  Tech Stack

* **Go** (Golang)
* **Chi Router** (`github.com/go-chi/chi`)
* **SQLite** (`modernc.org/sqlite` or `mattn/go-sqlite3`)
* **database/sql**

---

##  Project Structure

```text
.
├── db/            # database initialization
├── handlers/      # HTTP handlers
├── services/      # business logic
├── models/        # data structures
├── routes/        # router setup
├── main.go        # entry point
└── todos.db       # SQLite database
```

---

##  Setup & Run

### 1. Clone the repository

```bash
git clone https://github.com/arminwach/todo-api-go.git
cd todo-api-go
```

---

### 2. Install dependencies

```bash
go mod tidy
```

---

### 3. Run the server

```bash
go run main.go
```

---

### 4. Server will start on:

```text
http://localhost:8080
```

---

## 📡 API Endpoints

###  Create Todo

```http
POST /todos
```

```json
{
  "task": "Learn Go"
}
```

---

### 📄 Get All Todos

```http
GET /todos
```

---

###  Get Todo by ID

```http
GET /todos/{id}
```

---

###  Update Todo

```http
PUT /todos/{id}
```

```json
{
  "task": "Updated task"
}
```

---

###  Delete Todo

```http
DELETE /todos/{id}
```

---

##  Testing the API

You can test the API using:

* Thunder Client (VS Code)
* Postman
* curl

Example:

```bash
curl -X GET http://localhost:8080/todos
```

---

##  What I Learned

* How to build a REST API in Go
* How to structure backend projects
* How to interact with a database using `database/sql`
* How CRUD operations map to HTTP methods
* Debugging real backend issues (DB vs memory bugs)

---

##  Future Improvements

* Pagination (`GET /todos?page=1&limit=10`)
* Search/filtering
* Authentication (JWT)
* Logging & middleware
* Environment configuration
* Docker support

---

##  License

This project is licensed under the MIT License.

---

##  Author

**Armin Wach**
GitHub: https://github.com/arminwach
