all: *.go
	go build
	
test: *.go
	go test

p1: *.go
	go run . 17 5 15 3 2:35
	
p2: *.go
	go run . -w 17 5 15 3 2:35
	
p3: *.go
	go run . -v 17 5 15 3 2:35
	
p4: *.go
	go run . -v -w 17 5 15 3 2:35
