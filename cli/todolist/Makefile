BINARY_NAME=todolist

all: build run clean

build:
	go build -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
