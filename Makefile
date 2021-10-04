.PHONY: clean
.DEFAULT_GOAL := all

all: test build run

test:
	@go test .

build:
	@go build -mod=vendor .

run:
	@GO_EO_API_PORT=8081 \
	GO_EO_AUTHTOKEN=randomtokenforapi \
	./go-esperanto

clean:
	rm -rf ./go-esperanto
