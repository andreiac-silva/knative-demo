##Build Image
FROM golang:1.14-stretch AS builder
COPY . /invoice
WORKDIR /invoice
RUN go mod download && go get -u github.com/google/wire/cmd/wire && wire && CGO_ENABLED=0 GOOS=linux go build -o bin/application

##Run Image
FROM scratch
COPY --from=builder /invoice/bin/application application
EXPOSE 8080
ENTRYPOINT ["./application"]