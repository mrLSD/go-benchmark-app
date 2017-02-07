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
	@go test . -v -covermode=count
	@go vet -v  $(go list ./... 2>&1 | grep -v "vendor")
	@echo $(gocover)

build:
	@go build

fmt:
	@gofmt -w -l -s .

cover:
	@go test ./tools -covermode=count -coverprofile=c.out && go tool cover -html=c.out && unlink c.out

allcoverprofiles:
	@rm -rf *.coverprofile
	@sh ./test_cover.sh
	@gocovmerge *.coverprofile > merged.coverprofile
	@go tool cover -html=merged.coverprofile
	@rm -rf *.coverprofile
