.PHONY = all clean

SRCS := $(wildcard cmd/*.go)
BINS := $(SRCS:%.go=bin/%)

all: ${BINS}

bin/%: %.go
	@echo "Compiling executables..."
	go build -o $@ $<

clean: 
	@echo "Cleaning up..."
	rm -rf ${BINS}
