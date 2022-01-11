FROM golang:1.17 as builder

WORKDIR $GOPATH/src/composure/Hub

COPY . $GOPATH/src/composure/Hub

RUN go mod init
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Hub .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/composure/Hub/Hub .
CMD ["./Hub"]