fmt:
	find . -name "*.go" -exec go fmt {} \;

build:
	go build -o houston

clean:
	rm -f houston

.PHONY: fmt
