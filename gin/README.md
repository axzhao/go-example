docker build -t demo .
docker images
docker run -p 8080:8080 demo-docker

docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -v /data/docker-mysql:/var/lib/mysql -d mysql
docker run --link mysql:mysql -p 8080:8080 demo-docker

使用 Scratch 镜像

```
FROM scratch

WORKDIR $GOPATH/src/toki/go-demo/gin
COPY . $GOPATH/src/toki/go-demo/gin

EXPOSE 8000
CMD ["./gin"]
```

编译可执行文件
`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-demo .`
