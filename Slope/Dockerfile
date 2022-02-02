FROM golang:1.17 as builder

WORKDIR $GOPATH/src/composure/Slope

COPY . $GOPATH/src/composure/Slope

RUN go mod init
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Slope .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/composure/Slope/Slope .
COPY --from=builder /go/src/composure/Slope/config.json .
CMD ["./Slope"]