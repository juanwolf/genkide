.PHONY: test
test:
	go get -t && go test ./main.go

.PHONY: build
build:
	go build -o genkide .
