# docker build -t workflow-server -f deploy/build/Dockerfile .
# 第一个阶段：构建应用程序
FROM golang:1.18 AS builder

# 设置工作目录
WORKDIR /app

# 将本地代码复制到镜像中
COPY . .

# 构建Go应用
RUN go build -x -o server ./cmd

# 第二个阶段：创建最终镜像
# FROM alpine:latest
FROM debian:stable-slim

# 设置工作目录
WORKDIR /app

# 从第一个阶段复制构建产物
COPY --from=builder /app/server ./server

# 将配置文件复制到卷中
VOLUME /data/conf
COPY ./configs /data/conf

EXPOSE 8080

# 设置容器启动命令
CMD ["./server","-conf","/data/conf"]
