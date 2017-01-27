#
# Makefile
# @author Evgeny Ukhanov <mrlsd@ya.ru>
#

.PHONY: run, test, build, fmt
default: run

run:
	@go build && ./go-benchmark-app
	@rm -f ./go-benchmark-app

test:
	@go test . -covermode=count
	@go vet -v  $(go list ./... 2>&1 | grep -v "vendor")
	@echo $(gocover)

build:
	@go build

fmt:
	@gofmt -w -l -s .

cover:
	@go test . -tags all -coverprofile=c.out && go tool cover -html=c.out && unlink c.out