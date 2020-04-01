# LD_LIBRARY_PATH=$(LD_LIBRARY_PATH):.
CC=gcc
CFLAGS=-Wall
BIN=main
LIB=logger

all: $(BIN)

lib$(LIB).so lib$(LIB).h: liblogger.go
	go build -o lib$(LIB).so -buildmode=c-shared lib$(LIB).go

$(BIN): lib$(LIB).so lib$(LIB).h
	$(CC) $(CFLAGS) -o $(BIN) main.c -l $(LIB) -L.

run: $(BIN)
	./$(BIN)
