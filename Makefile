## easy: generate easyjson
easy:
	cd internal/pkg/responses && easyjson --all responses.go
	cd .. && cd .. && cd ..
	cd internal/app/task/models && easyjson --all task.go

## build: build container
build:
	docker build --no-cache -t gotbit .

## run: run container
run:
	docker run -d -p 8081:8081 -p 8080:8080 gotbit
	echo "Documentaion on http://localhost:8080/swagger/index.html"

## local: local run
local:
	go run cmd/main.go

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command to run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo