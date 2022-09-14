.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

## ------------------------------------------------- Common commands: --------------------------------------------------
## Formats the code.
format:
	${call colored,formatting is running...}
	go vet ./...
	go fmt ./...

## Fix-imports order.
fix-imports:
	${call colored,fixing imports...}
	./scripts/fix-imports-order.sh
