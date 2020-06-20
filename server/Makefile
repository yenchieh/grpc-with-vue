.PHONY: deps vet test dev build build-frontend build-image clean

DOCKER_REPO_URL = vote

deps:
	go mod download && go mod vendor

dev:
	go run main.go

build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ./build/main *.go

push-image: build build-frontend build-image
	docker push $(DOCKER_REPO_URL):latest
