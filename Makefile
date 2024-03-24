build:
	go build -o bin/openalex-cli main.go

install: build
	go install

create-tables:
	openalex-cli create-tables

drop-tables:
	openalex-cli drop-tables

create-indexes:
	openalex-cli create-indexes

drop-indexes:
	openalex-cli drop-indexes
