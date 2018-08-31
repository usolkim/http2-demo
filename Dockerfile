FROM golang:1.11-rc
COPY . /go/src/app

WORKDIR /go/src/app/src

EXPOSE 80
EXPOSE 443
RUN go build -o http2-demo main.go
CMD ["./http2-demo"]