export AOC_COOKIE=

all: run

run:
	DAY=1 go run cmd/aocgo2025/main.go

test:
	go test ./...

new:
	mkdir internal/day
	cp internal/template/template.go internal/day/day.go
	cp internal/template/template_test.go internal/day/day_test.go

.SILENT: run test new
.PHONY: run test new
