FROM golang:1.20.7-alpine3.18 AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 下载依赖信息
RUN go mod download

# 将我们的代码编译成二进制可执行文件 tiktok
RUN go build -o tiktok .

FROM python:3.11.5-slim-bookworm

MAINTAINER hiifong<i@hiif.ong>

RUN pip install coversnap
RUN pip uninstall -y opencv-python
RUN pip install opencv-python-headless

WORKDIR /app

COPY ./conf/tiktok.yml /app/conf/

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/tiktok /app

EXPOSE 8080

# 需要运行的命令
ENTRYPOINT ["/app/tiktok", "server"]
