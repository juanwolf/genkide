.PHONY: build
build:
	go build -o genkide .

.PHONY: test
test:
	go get -t && go test ./main.go
