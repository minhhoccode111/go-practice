.PHONY: all
all: build

.PHONY: build
build:
	go build -o ./bin/out

.PHONY: run
run: build
	./bin/out

.PHONY: clean
clean:
	rm -fr ./bin/
