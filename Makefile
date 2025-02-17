dev:
	go run cmd/main.go
startDB:
	docker-compose up -d
stopDB:
	docker-compose down
migrate:
	go run migration/main.go			