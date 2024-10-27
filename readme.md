# Student management

## Description

This project is a RESTful API for managing student data, built using Golang and SQLite. It supports CRUD (Create, Read, Update, Delete) operations for handling student information, allowing the following functionalities:

- Manage students with CRUD operations
- Get list of students
- Get student by id
- Update student by id
- Delete student by id

## Dependencies
- Golang: A statically typed programming language designed for building scalable and efficient applications.
- Sqlite: A lightweight, serverless database engine stored in a single file, suitable for applications with limited data storage needs.

## Configuration
- ```env```: Environment configuration for development and production setups.
- ```storage_path```: File path for SQLite database storage.
- ```http_server```: HTTP server configuration for setting the server’s listening address and port.

Configuration parameters are set in a YAML file (e.g., ```config/local.yml```).


## sample configuration (config/local.yml)
```yaml
    env: "dev"
    storage_path: "storage/storage.db"
    http_server:
      address: "localhost:8085"
```

## Getting started

### Prerequisites
- ```Go``` installed on the machine
- clone the repository
- Ensure the ```config/local.yml``` file is present in the ```config``` folder or setup correctly.

Run the following command to start the server 
```bash
    go run cmd/students-management/main.go -config config/local.yml 
```

### Output
```bash
    2024/10/27 14:17:46 INFO Server started address=localhost:8085
```

## API Endpoints
### Base URL
```bash
    http://localhost:8085/api/v1
```
1. Create a new student

```bash
    POST http://localhost:8085/api/v1/students
    header: Content-Type: application/json
    body: {
        "name": "Alice Johnson",
        "age": 19,
        "email": "alice@example.com",
    }
```
sample response: 
```json
    {
        "id": 1
    }
```
2. Get list of students

```bash
    GET http://localhost:8085/api/v1/students
```
sample response: 
```json
[
    {
        "id": 1,
        "name": "sahil",
        "email": "sahil@dev.com",
        "age": 24
    },
    {
        "id": 4,
        "name": "kshitij",
        "email": "kshitij@dev.com",
        "age": 26
    }
]
```
3. Get student by id

```bash
    GET http://localhost:8085/api/v1/students/4
```
sample response: 
```json
{
    "id": 4,
    "name": "kshitij",
    "email": "kshitij@dev.com",
    "age": 26
}
```

4. Update student by id

```bash
    PATCH http://localhost:8085/api/v1/students/4
    header: Content-Type: application/json
    body: {
        "name": "kshitij",  # update name
        "email": "kshitij@dev.com",  # update email
        "age": 26  # update age
    }
```
sample response: 
```json
{
    "id": 4,
    "name": "kshitij",
    "email": "kshitij@dev.com",
    "age": 26
}
```
5. Delete student by id

```bash 
    DELETE http://localhost:8085/api/v1/students/4
```     
sample response: 
```json
{
    "message": "Student deleted successfully."
}
```

## Testing with postman
To facilitate API testing, you can import these request examples into Postman:

- Open Postman and create a new collection.
- Add requests for each endpoint, copying the details from the examples above.
- Run the requests individually or as a collection to interact with the API.