FROM golang:1.17 as builder

WORKDIR $GOPATH/src/composure/Balance

COPY . $GOPATH/src/composure/Balance

RUN go mod init
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Balance .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/composure/Balance/Balance .
CMD ["./Balance"]