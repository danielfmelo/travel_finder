img := danielfmelo/travel_finder:latest
wd := $(shell pwd)
cachevol=$(wd)/.gomodcachedir:/go/pkg/mod
rundocker := docker run --rm -v $(wd):/app -v $(cachevol) $(img)
database?=

image:
	docker build . -t $(img)

run: 
	go run cmd/travel_finder.go $(database)

docker-run: image docker-build
	./travel_finder $(database)

build:
	go build -o ./travel_finder ./cmd/travel_finder.go

docker-build: image
	$(rundocker) go build -v -o ./travel_finder ./cmd/travel_finder.go

unit-tests: 
	go test -timeout 20s -tags unit -race -coverprofile=coverage.out ./...

docker-tests: image
	$(rundocker) go test -timeout 20s -tags unit -race -coverprofile=coverage.out ./...

coverage: unit-tests
	go tool cover -html=coverage.out -o=coverage.html
	xdg-open coverage.html
