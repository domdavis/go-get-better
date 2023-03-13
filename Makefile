all: clean build test vet lint

run:
	go run solution.go

clean:
	go clean
	rm -f *.out *.prof *.test
	go mod tidy

build:
	go build -o solution

vet:
	go vet ./...

lint:
	golangci-lint run
	golint ./...

test:
	go test -v -covermode=count ./...
