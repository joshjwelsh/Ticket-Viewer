dev:
	go run *.go

build:
	go build . 

build+:
	go build . && ./viewer
