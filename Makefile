.PHONY: test build run

test:
	go test ./...

build:
	go build -o streakforge ./cmd/streakforge

run:
	go run ./cmd/streakforge help

