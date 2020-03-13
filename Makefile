clean:
	rm a.out

build:
	go build -o a.out .

run: build
	./a.out