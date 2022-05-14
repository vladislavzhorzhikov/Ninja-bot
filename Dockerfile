FROM 192.168.0.8:5000/golang

EXPOSE 80

WORKDIR /go/src/Ninja-bot
COPY . .
RUN chmod +x ./start.sh

RUN go install -mod vendor

ENTRYPOINT ./start.sh
