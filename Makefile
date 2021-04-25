all: test

PROCS = 2

help:                        ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-20s%s\n", $$1, $$2}'

init:                        ## Install tools.
	go generate -x -tags=tools ./tools
	make build

build:                       ## Build with race detector.
	go install -v -race ./...

test: build                  ## Test with race detector.
	go test -v -race ./...

fuzz-reverse:                ## Fuzz reverse function using dev.fuzz.
	go test -v -race -fuzz=FuzzReverse -fuzztime=50000x -parallel=$(PROCS)

fuzz-protocol:               ## Fuzz protocol using dev.fuzz.
	cd protocol && go test -v -race -fuzz=FuzzHandler -parallel=$(PROCS) -timeout=10s

gofuzz-protocol:             ## Fuzz protocol using dvyukov/go-fuzz.
	cd protocol && ../bin/go-fuzz-build -race
	cd protocol && ../bin/go-fuzz -procs=$(PROCS)

gofuzz-reverse:              ## Fuzz reverse function using dvyukov/go-fuzz.
	./bin/go-fuzz-build -race
	./bin/go-fuzz -procs=$(PROCS)
