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

coverprofiles:
	@rm -rf *.coverprofile
	@for Package in `go list ./... | grep -v "vendor"` ; do \
		x=`cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 5 | head -n 1` ; \
		echo $$(go test -covermode=count -coverprofile=$$x.coverprofile $$Package) ; \
		echo $$Package ; \
	done

