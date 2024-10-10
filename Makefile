# app name
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs
