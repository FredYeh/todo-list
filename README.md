## Coding Exercise

RESTful api for todo-list

- GET /tasks
- POST /tasks
- PUT /tasks/{id}
- DELETE /tasks/{id}
----
```go
type Task struct {
    Name string
    Status int //0 for incompleted, 1 for completed
}
```
-----
- Go version: 1.19.3
- Database: redis 7.2.4
