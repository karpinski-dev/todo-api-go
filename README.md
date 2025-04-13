## How to start the server
To start this simple API run `go run .`. The server runs at `localhost:8080`.
## Endpoints
1. `/all` - returns all todo items
2. `/done` - returns items with status `true`
3. `/pending` - returns items with status `false`
## Data
This repository has static data with the following data structure:
```go
type TodoItem struct {
	Title   string `json:"title"`
	Details string `json:"details"`
	Status  bool   `json:"status"`
}
```
