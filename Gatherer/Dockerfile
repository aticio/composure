FROM golang:1.17 as builder

WORKDIR $GOPATH/src/composure/Gatherer

COPY . $GOPATH/src/composure/Gatherer

RUN go mod init
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Gatherer .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/composure/Gatherer/Gatherer .
CMD ["./Gatherer"]