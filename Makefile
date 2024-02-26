build:
	go build -o server main.go

run: build
	./server

clean:
	rm -f server

watch:
	reflex -s -r '\.go$$' make run