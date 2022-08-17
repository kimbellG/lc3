run: build
	./lc3 ~/Downloads/2048.obj 2> /dev/null
debug: build
	./lc3 ~/Downloads/2048.obj 2> logs
build:
	go build -o lc3 cmd/lc3/lc3.go
	
