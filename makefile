dev:
	go run auth.go accessor.go const.go model.go app.go display.go device.go

build:
	go build . 

run:
	go build . && ./main

test:
	go test -cover *.go