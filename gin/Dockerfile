FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/toki/go-example/gin
COPY . $GOPATH/src/go-example/gin
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./gin"]