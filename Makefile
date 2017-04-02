# Makefile

usage:
	@echo ""
	@echo "Task                 : Description"
	@echo "-----------------    : -------------------"
	@echo "make setup           : Install all necessary dependencies"
	@echo "make build           : Generate production build for current OS"
	@echo "make test            : Run tests"
	@echo "make restore-vendors : Fetch dependencies to the vendor / directory"
	@echo ""

build:
	gb build

test:
	gb test

serve:
	gb build
	bin/gsm

dev-deps:
	go get github.com/constabulary/gb/...

restore-vendors:
	gb vendor restore

setup: dev-deps restore-vendors
