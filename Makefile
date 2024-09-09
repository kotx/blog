.PHONY: build generate

build: generate
	go run main.go

generate:
	go install github.com/a-h/templ/cmd/templ@latest
	go generate
