.PHONY = all clean

SRCS := $(wildcard cmd/*.go)
BINS := $(SRCS:cmd/%.go=bin/%)

V_SRCS := $(wildcard cmd/*.v)
V_BINS := $(V_SRCS:cmd/%.v=bin/%)

all: ${BINS} ${V_BINS}

bin/%: cmd/%.go
	@echo "Compiling executables..."
	go build -o $@ $<

bin/%: cmd/%.v
	@echo "Compiling executables..."
	v -o $@ $<

clean: 
	@echo "Cleaning up..."
	rm ${BINS} ${V_BINS}
