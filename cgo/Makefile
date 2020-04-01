CC=gcc
CFLAGS=-Wall

all: main liblogger.so liblogger.h

clean:
	rm main liblogger.h liblogger.so

liblogger.so liblogger.h: liblogger.go
	go build -o liblogger.so -buildmode=c-shared liblogger.go

main: liblogger.so liblogger.h
	$(CC) $(CFLAGS) -o main main.c -l logger -L.
