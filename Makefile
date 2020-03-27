clean:
	rm a.out

build:
	go build -o a.out ./cmd

run: build
	./a.out
