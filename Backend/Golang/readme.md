### Project Structure

```md
project
├── cmd
│ ├── main.go ## enter point of the project
├── config
│ ├── app.go ## database configurations
├── internal
│ ├── dto
│ ├── entities
│ ├── handler ## handler request (controller)
│ ├── model
│ ├── repository
│ └── service ## business logic
├── pkg
│ ├──fake_data.go
│ ├──middleware.go
│ ├──router.go
└── readme.md
```

## Run test

Test required integration test

inside internal folder move to handler and run

```go
go run test .
```
