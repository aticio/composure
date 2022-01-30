FROM golang:1.17 as builder

WORKDIR $GOPATH/src/composure/Dealer

COPY . $GOPATH/src/composure/Dealer

RUN go mod init
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Dealer .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/composure/Dealer/Dealer .
CMD ["./Dealer"]