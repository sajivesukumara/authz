# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=authz
BINARY_UNIX=$(BINARY_NAME)_unix

# All target
all: test build

# Test target
unittest:
	$(GOTEST) -v ./...

init:
	$(GOCMD) mod init $(BINARY_NAME)
	$(GOCMD) get github.com/golang-jwt/jwt/v5

# Build target
build:
	$(GOCMD) mod tidy
	$(GOCMD) get -v ./...
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/

# Migration build target
mbuild:
	$(GOCMD) mod tidy
	$(GOCMD) get -v ./...
	$(GOBUILD) -o createdb ./cmd/migrate
	./createdb

# Clean target
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run target
run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/
	./$(BINARY_NAME)

# Run curl tests
#curl --header "Content-Type: application/json"   --request POST   --data "{\"username\":\"admin\",\"password\":\"admin\"}"   http://localhost:8000/login   
#curl --header "Content-Type: application/json"   --request GET   --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkyMzcxMjAsInVzZXJuYW1lIjoiYWRtaW4ifQ.KCpE3Ey0CpT-MfAZtLhj74HncrWmrYRvI8wvUvbph8A "  http://localhost:8000/protected  

test:
	
	curl --header "Content-Type: application/json" --request POST    http://localhost:8000/auth/signup --data "{\"username\":\"admin\",\"password\":\"admin\"}"

# Docker target
docker:
	docker build -t $(BINARY_NAME) .
