package = github.com/rgoomar/drun

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/drun-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/drun-linux-386 $(package)
	GOOS=linux GOARCH=arm go build -o release/drun-linux-arm $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/drun-darwin-amd64 $(package)
	GOOS=darwin GOARCH=386 go build -o release/drun-darwin-386 $(package)

.PHONY: release
