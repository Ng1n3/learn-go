.DEFAULT_GOAL := run

.PHONY:fmt vet build run
fmt:
	go fmt ./main.go
vet: fmt
	go vet ./main.go
build: vet
	go build -o variadic ./main.go
run: build
	./variadic