.PHONY: all
all: generate css

.PHONY: generate
generate:
templ generate ./src/...
mkdir -p dist
cp -r src/components dist/
find dist -name "*.templ" -type f -delete

.PHONY: css
css:
npm run css:build

.PHONY: watch
watch:
air

.PHONY: dev
dev:
make generate
make css
make watch

.PHONY: clean
clean:
rm -rf dist
