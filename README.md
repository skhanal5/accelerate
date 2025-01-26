# Accelerate

## About
Quick little project to brush up on Go.

## Local Development

### Running Locally
Start up the server using `go run cmd/accelerate/main.go`

### Running Unit Tests
Project convention: Each unit test is defined alongside the source file in the same package.

To run all unit tests, use the following command:
```bash
go test ./...
```

To run all unit tests with coverage: 
```bash
# with coverage
go test --cover ./...

# export coverage to an html
go test ./... -coverprofile=coverage_report
go tool cover -html=coverage_report
```