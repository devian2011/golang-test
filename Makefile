all: test build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main/main.go

test:
	go test -v ./...
