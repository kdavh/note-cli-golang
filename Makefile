install:
	# builds and install binary to $GOPATH/bin/note-cli-golang
	# use `go build` to instead output binary to cwd
	go install

update:
	# if updates happened elsewhere, use git to manually pull changes, then `make install`
	# `git pull && go install`
	# or just have go pull and build:
	go install github.com/kdavh/note-cli-golang
