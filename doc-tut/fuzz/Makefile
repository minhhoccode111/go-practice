.PHONY: build run

build:
	go build -o bin/app .

run: build
	./bin/app

clean:
	rm -f bin/app

test:
	go test

test-fuzz:
	go test -fuzz=Fuzz
