##Build Image
FROM golang:1.14-stretch AS builder
COPY . /payment
WORKDIR /payment
RUN go mod download && go get -u github.com/google/wire/cmd/wire && wire && CGO_ENABLED=0 GOOS=linux go build -o bin/application

##Run Image
FROM scratch
COPY --from=builder /payment/bin/application application
EXPOSE 9001
ENTRYPOINT ["./application"]