# Number Classification API

This API classifies a number based on its mathematical properties, such as whether it is prime, perfect, or armstrong. It also calculates the digit sum and provides a fun fact about the number from the Numbers API.

---

## Project Description

The **Number Classification API** allows users to query for the classification properties of a number. The response includes details such as:
- Prime status
- Perfect number status
- List of properties (like odd, even, or Armstrong)
- Sum of the digits
- A fun fact about the number from [Numbers API](http://numbersapi.com/).

The API is built using the **Gin** framework and follows a simple RESTful structure.

---

## Instructions to Run the API Locally

### Prerequisites
- Install [Go](https://go.dev/doc/install) (version 1.19 or higher).
- Ensure you have `git` installed.

### Steps to Run Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/Jidetireni/num_class_api.git
   cd num_class_api.git
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the server:
   ```bash
   go run main.go
   ```

4. The server will be running at:
   ```bash
   http://localhost:8080
   ```

---

## Deployment Details

To deploy this API, follow these general steps:

### On a Cloud Platform (e.g., AWS EC2)
- Ensure the server has Go installed.
- Copy the project files to the server.
- Build the Go binary:
  ```bash
  go build -o num_class_api main.go
  ```
- Run the binary:
  ```bash
  ./num_class_api
  ```
- Ensure the appropriate ports are opened for inbound traffic.


```
 .
docker run -p 8080:8080 num_class_api
```

---

## Example API Requests and Responses

### Request
```
GET /api/classify-number?number=28
```
### Response
```json
{
  "number": 28,
  "is_prime": false,
  "is_perfect": true,
  "properties": ["even"],
  "digit_sum": 10,
  "fun_fact": "28 is a perfect number."
}
```
### Error Request
```
GET /api/classify-number?number=abc
```
### Error Response
```json
{
  "number": "alphabet",
  "error": true
}
```

---

## Directory Structure
```
.
├── main.go                  # Entry point of the application
├── handlers                 # API request handlers
│   └── classify.go          # Number classification logic
├── utils                    # Utility functions
│   └── helper.go            # Helper functions for number properties
├── go.mod                   # Go module file
├── go.sum                   # Dependency checksum
└── .gitignore               # Git ignore file
```

---

## Improvements & Notes
- Add more number classifications (like Fibonacci check).
- Handle external API errors more gracefully.
- Use logging for better debugging.

---
