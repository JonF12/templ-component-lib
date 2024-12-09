all: generate css

build:
	@go build -o ./main ./examples/server/main.go

generate:
	@templ generate ./src/**/*
	@mkdir -p dist
	@cp -r src/ dist/
	@find dist -name "*.templ" -type f -delete

run:
	@go run examples/server/main.go

css:
	@npm run css:build

watch:
	@air

dev:
	@make generate
	@make build
	@make css
	@make run

clean:
	@rm -rf dist
