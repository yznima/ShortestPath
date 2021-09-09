.PHONY: build clean

build: 
	go build -v -o shortest-path *.go 

clean: 
	rm shortest-path