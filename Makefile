run:
	go run cmd/main.go

build:
	docker build -t to_do:latest .

run-compose:
	docker-compose up