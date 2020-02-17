.PHONY: test clean

build:
	mkdir -p functions
	go get ./...
	go build -o ./functions/ ./...

netlify:
	mkdir -p functions
	go get ./...
	go install ./...

test:
	go test ./...

clean:
	rm -f functions/*
