# LD_LIBRARY_PATH=$(LD_LIBRARY_PATH):.

build: liblogger.go
	go build -o liblogger.so -buildmode=c-shared liblogger.go

link: liblogger.so liblogger.h
	gcc -Wall -o main main.c -l logger -L.

run: link 
	./main
