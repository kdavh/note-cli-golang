PHONY: i
# for local development, just run this and then reference new file in bin
# builds and install binary to $GOPATH/bin/note-cli-golang
# use `go build` to instead output binary to cwd
install:
	go install
i: install

# if updates happened elsewhere, this gets those updates
# while keeping other things, like custom local git remotes, working changes
# basically it's a fancy git pull
# same thing from this dir, but less educational: `go get -u`
update:
	go get -u github.com/kdavh/note-cli-golang

fix-git:
	# golang automatically pulls using https, add ssh:
	git remote add originssh git@github.com:kdavh/note-cli-golang.git
	g fetch originssh master
	g branch --set-upstream-to=originssh/master
