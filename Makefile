local: clean
	go build
binary: clean
	env GOOS=linux GOARCH=amd64 go build -v
lambda: binary
	zip membership-salesforce-keepalive.zip main.js membership-salesforce-keepalive
clean:
	rm membership-salesforce-keepalive.zip membership-salesforce-keepalive || exit 0
deploy: all
	aws s3 cp membership-salesforce-keepalive.zip s3://membership-dist/membership/PROD/membership-salesforce-keepalive/membership-salesforce-keepalive.zip --profile membership
all: clean lambda
run: local
	env AWS_PROFILE=membership ./membership-salesforce-keepalive
