fmt:
	find . -name "*.go" -exec go fmt {} \;

build:
	go build -o houston main.go

clean:
	rm -f houston

.PHONY: fmt
