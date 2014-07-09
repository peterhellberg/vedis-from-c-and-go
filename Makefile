all: build run

build:
	@echo "Building the examples:"
	go build -o bin/example-from-go example.go
	gcc example.c vedis.c -o bin/example-from-c

run:
	@echo "\nRunning C example:\n"
	@./bin/example-from-c data.vdb
	@echo "\nRunning Go example:\n"
	@./bin/example-from-go data.vdb

clean:
	@echo "\nCleanup:"
	rm -f bin/example-from-c bin/example-from-go
	rm -f data.vdb
