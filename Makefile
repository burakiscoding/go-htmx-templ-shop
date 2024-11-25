BINARY_NAME=main

gen_templ:
	templ generate

build:
	templ generate
	go build -o bin/${BINARY_NAME} main.go

run: build
	bin/${BINARY_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}