
---

##  Library Management API

### Description

A Go Fiber-based API backend for managing a library system. It allows users to manage books, authors, and borrowers. Users can borrow books and track who has borrowed which books.

### Requirements

- **Golang 1.23+**



---
### Dependency

- **Fiber Framework**
- **SQLite Database**
- **GORM ORM**

---

### Build and Run the Application

1. **Clone the repository**:
   ```bash
   git clone https://github.com/progressive-galib/library-management-api-jwt.git
   cd library-management-api-jwt
   ```

2. **Get the Dependency**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

   The application will start at `http://localhost:3000`.
---

### Running Automated Tests

To run tests for the application:

```bash
./gradlew test
```

---


### API Routes



#### 1. **Books Management**

 **POST /api/v1/books**  
_Adds a new book to the library._

**Payload**:  
```json
{
  "title": "string",
  "author_id": "int",
  "available": "boolean"
}
```

**Example Use**:  
```bash
curl -X POST http://localhost:3000/api/v1/books \
-H "Content-Type: application/json" \
-d '{
  "title": "The Go Programming Language",
  "author_id": 1,
  "available": true
}'
```

---

 **GET /api/v1/books**  
_Lists all books in the library._

**Response**:  
```json
[
  {
    "id": "int",
    "title": "string",
    "author_id": "int",
    "available": "boolean"
  }
]
```

**Example Use**:  
```bash
curl -X GET http://localhost:3000/api/v1/books
```

---

#### 2. **Author Management**

 **POST /api/v1/authors**  
_Adds a new author._

**Payload**:  
```json
{
  "name": "string"
}
```

**Example Use**:  
```bash
curl -X POST http://localhost:3000/api/v1/authors \
-H "Content-Type: application/json" \
-d '{
  "name": "Alan A. A. Donovan"
}'
```

---

 **GET /api/v1/authors**  
_Lists all authors._

**Response**:  
```json
[
  {
    "id": "int",
    "name": "string"
  }
]
```

**Example Use**:  
```bash
curl -X GET http://localhost:3000/api/v1/authors
```

---

#### 3. **Borrowing Management**

 **POST /api/v1/borrow**  
_Borrow a book for a user._  

**Payload**:  
```json
{
  "book_id": "int",
  "user_id": "int"
}
```

**Example Use**:  
```bash
curl -X POST http://localhost:3000/api/v1/borrow \
-H "Content-Type: application/json" \
-d '{
  "book_id": 1,
  "user_id": 1
}'
```

---

 **GET /api/v1/borrowed/:user_id**  
_Lists all books borrowed by a specific user._  

**Request Parameters**:  
- `user_id`: The ID of the user.

**Response**:  
```json
[
  {
    "id": "int",
    "book_id": "int",
    "user_id": "int",
    "due_date": "string (ISO 8601)"
  }
]
```

**Example Use**:  
```bash
curl -X GET http://localhost:3000/api/v1/borrowed/1
```

---

### Summary of All Routes

| Route                         | Method | Description                               |
|-------------------------------|--------|-------------------------------------------|
| `/api/v1/books`               | POST   | Adds a new book                           |
| `/api/v1/books`               | GET    | Lists all books                           |
| `/api/v1/borrow`              | POST   | Borrow a book                             |
| `/api/v1/borrowed/:user_id`   | GET    | Lists all books borrowed by a specific user |
| `/api/v1/authors`             | POST   | Adds a new author                         |
| `/api/v1/authors`             | GET    | Lists all authors                         |

---


## Future Enhancements

- Add user authentication with JWT for better security.
- Add functionality for returning books and marking them as available.
- Implement pagination and filtering for listing books and authors.

---

## License

This project is licensed under the MIT License.

```

This `README.md` should give any developer an understanding of the project, its purpose, and how to get started with it. Adjust any placeholders like `your-username` with the actual details, and add any additional notes specific to your setup if needed!
