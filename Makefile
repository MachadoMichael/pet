build:# Define the name of the binary
BINARY_NAME=./bin/pet

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

# Turn on database
db:
	@echo "Openning Docker Desktop..."
	open -a Docker &
	sleep 10
	@echo "Docker Desktop opened. Running docker compose"
	docker compose up -d
	@echo "Docker compose started. Running migrations"
	sleep 3
	$(BINARY_NAME) migrations

# Run the binary
run: build
	$(BINARY_NAME)

migrations:
	$(BINARY_NAME) migrations
# Clean up the binary
clean:
	rm -f $(BINARY_NAME)

# run tests
test:
	go test -v ./...
