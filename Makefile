
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/linux/amd64/drone-docker-compose ./;

.PHONY: build