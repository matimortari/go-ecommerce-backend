# Build the Go application
build:
	@go build -o bin/ecom cmd/main.go

# Run tests for the Go project
test:
	@go test -v ./...

# Run the application (depends on the build)
run: build
	@./bin/ecom

# Rule for creating migrations
migration-create:
	@migrate create -ext sql -dir cmd/migrate/migrations $(name)

# Run migrations "up"
migrate-up:
	@go run cmd/migrate/main.go up

# Run migrations "down"
migrate-down:
	@go run cmd/migrate/main.go down

migrate-status:
	@go run cmd/migrate/main.go version
