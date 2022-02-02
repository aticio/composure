FROM golang:1.17 as builder

WORKDIR $GOPATH/src/composure/Pearson

COPY . $GOPATH/src/composure/Pearson

RUN go mod init
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Pearson .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/composure/Pearson/Pearson .
COPY --from=builder /go/src/composure/Pearson/config.json .
CMD ["./Pearson"]