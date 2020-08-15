##Build Image
FROM golang:1.14-stretch AS builder
COPY . /email
WORKDIR /email
RUN go mod download && go get -u github.com/google/wire/cmd/wire && wire && CGO_ENABLED=0 GOOS=linux go build -o bin/application

##Run Image
FROM scratch
COPY --from=builder /email/bin/application application
EXPOSE 9002
ENTRYPOINT ["./application"]