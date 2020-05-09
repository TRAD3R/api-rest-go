# Файл для указания консольных команд
# Команда build собирает бинарник в папку cmd/apiserver
.PHONY: build
build:
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build # команда build вызывается по умолчанию
