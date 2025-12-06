export AOC_COOKIE=

all: run

run:
	DAY=6 go run cmd/aocgo2025/main.go

test:
	go test ./...

.SILENT: run test
.PHONY: run test
