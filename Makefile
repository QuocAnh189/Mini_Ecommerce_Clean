swag:
	swag init -g internals/server/http/server.go

docker-compose up:
	docker-compose --env-file .env up --build -d 

docker-compose down:
	docker-compose --env-file .env down