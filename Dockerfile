FROM golang:1.20.5-alpine AS builder

# ENV GO111MODULE on
# ENV GOPROXY=https://goproxy.io,direct

RUN apk add --update --no-cache ca-certificates git

WORKDIR $GOPATH/src/github.com/kekasicoid/restapi-socketio

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o kekasigen app/main.go

# Start fresh from a smaller image
FROM alpine:3.9 

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /root

COPY --from=builder /go/src/github.com/kekasicoid/restapi-socketio/kekasigen .
COPY --from=builder /go/src/github.com/kekasicoid/restapi-socketio/.env .env

CMD [ "./kekasigen" ]