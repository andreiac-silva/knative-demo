SENSEDIA_APP_NAME = email-sender

build:
	go mod download
	go get -u github.com/gobuffalo/packr/v2/...
	packr2
	go get -u github.com/google/wire/cmd/wire
	wire
	CGO_ENABLED=0 GOOS=linux go build -o bin/application

image:
	cp /etc/ssl/certs/ca-certificates.crt .
	docker build -t purchase .