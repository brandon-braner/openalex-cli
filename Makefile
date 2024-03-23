build:
	go build -o bin/openalex-cli main.go

install: build
	go install