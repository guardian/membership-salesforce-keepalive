goos:
	export GOOS=linux
goarch:
	export GOARCH=amd64
binary: goos goarch
	go build
lambda: binary
	zip membership-salesforce-keepalive.zip main.js membership-salesforce-keepalive
clean:
	rm membership-salesforce-keepalive.zip membership-salesforce-keepalive
all: lambda
