.PHONY = all clean

LD_LIBRARY_PATH=.

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

lib/liblogger.so bin/liblogger.h: cgo/liblogger.go
	go build -o lib/liblogger.so -buildmode=c-shared $<

bin/cgo: cmd/cgo.c lib/liblogger.so lib/liblogger.h
	$(CC) $(CFLAGS) -o bin/cgo cmd/cgo.c -l logger -L./lib

clean: 
	@echo "Cleaning up..."
	rm ${BINS} ${V_BINS} lib/* bin/cgo
