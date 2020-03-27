OUT=a.out

all: build

clean:
	[ -f "$(OUT)" ] && rm $(OUT) || echo "Nothing to clean"

build:
	go build -o $(OUT) ./cmd

run: build
	./$(OUT)

