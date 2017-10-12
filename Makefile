prog: math.h math.c main.go
	go build -o prog
clean:
	rm -rf prog
	