install:
	# builds and install binary to $GOPATH/bin/note-cli-golang
	# use `go build` to instead output binary to cwd
	go install

update:
	# if updates happened elsewhere, this gets those updates
	# while keeping other things, like custom local git remotes
	go get -u github.com/kdavh/note-cli-golang

fix-git:
	# golang automatically pulls using https, switch to ssh:
	git remote set-url origin git@github.com:kdavh/note-cli-golang.git
