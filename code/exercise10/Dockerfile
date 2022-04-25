##源镜像
FROM golang:latest  AS builder
##作者
MAINTAINER "test@163.com"
#设置工作目录
WORKDIR $GOPATH/src/github.com/golangcode
#将服务器的go工程代码加入到docker容器中
ADD web_server.go $GOPATH/src/github.com/golangcode
#go构建可执行文件
RUN go build web_server.go

FROM scratch
COPY --from=builder   $GOPATH/src/github.com/golangcode /
#暴露端口
EXPOSE 8888
#最终运行docker的命令
ENTRYPOINT  ["./web_server"]
