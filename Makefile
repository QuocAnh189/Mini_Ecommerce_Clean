.PHONY: run swag docker-compose-up build-app build-frontend push-app push-frontend

run:
	go run cmd/app/main.go

build:
	GOOS=linux GOARCH=arm64 go build -o dist/ecommerce ./cmd/app

swag:
	swag init -g internals/server/http/server.go

docker-compose-up:
	docker-compose --env-file .env up --build -d 

docker-compose-down:
	docker-compose --env-file .env down

build-app:
	docker build -t anhquoc1809/ecommerce.app .

build-frontend:
	docker build -t anhquoc1809/ecommerce.frontend ./frontend

push-app:
	docker push anhquoc1809/ecommerce.app

push-frontend:
	docker push anhquoc1809/ecommerce.frontend

 