db-up:
	docker compose up -d

db-down:
	docker compose down

run:
	go run main.go
