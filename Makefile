BINARY_NAME=x

build: 
	@go build -C cmd -o ../bin/$(BINARY_NAME)

run: build
	@./bin/$(BINARY_NAME)

clean:
	@rm -f bin/$(BINARY_NAME)

.PHONY: build run clean
