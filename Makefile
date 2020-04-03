.PHONY = all clean

SRCS := $(wildcard cmd/*.go)
BINS := $(SRCS:cmd/%.go=bin/%)

all: ${BINS}

bin/%: cmd/%.go
	@echo "Compiling executables..."
	go build -o $@ $<

clean: 
	@echo "Cleaning up..."
	rm -rf ${BINS}
