# Makefile

usage:
	@echo ""
	@echo "Command              : Description"
	@echo "-------------------- : --------------------------------------------"
	@echo "make setup           : Install all necessary dependencies"
	@echo "make build           : Generate production build for current OS"
	@echo "make test            : Run tests"
	@echo "make serve           : Run the app locally"
	@echo "make image           : Create docker image"
	@echo "make run-image       : Run the app on docker"
	@echo "make clean           : Remove dependencies and compiled files"
	@echo ""

build:
	gb build

test:
	gb test

image:
	docker build -t husniadil/gsm .

run:
	bin/gsm

dev-deps:
	go get github.com/constabulary/gb/...

restore-vendors:
	gb vendor restore

setup: dev-deps restore-vendors

serve: build run

run-image:
	docker run -it --rm -p 8080:8080 husniadil/gsm

clean:
	rm -Rf vendor/src
	rm -Rf bin
	rm -Rf pkg
