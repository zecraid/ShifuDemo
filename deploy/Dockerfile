# 阶段1：拉取代码并构建应用
FROM golang:latest AS builder

WORKDIR /app

# 拉取代码
RUN git clone https://github.com/zecraid/ShifuDemo.git .

# 构建应用
RUN go build -o main ./src/test.go

# 阶段2：创建最终的镜像
FROM alpine:latest

WORKDIR /app

# 从第一个阶段拷贝构建好的应用
COPY --from=builder /app/main .

# 启动应用
CMD ["./main"]
