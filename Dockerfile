FROM golang:alpine

WORKDIR /go/src/app

ADD . .
RUN go mod init

RUN go build  -o /Ninja-bot

EXPOSE 6111

CMD ["./Ninja-bot"]