BINARY_NAME=split-file-to-chunks

all: build run clean

build:
	go build -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
	# remove every other md files except README
	find . -type f -name "*.md" ! -name "README.md" -delete
