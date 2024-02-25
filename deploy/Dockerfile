# 阶段1：拉取代码并构建应用
FROM golang:latest AS builder

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

# 拉取代码
RUN git clone https://github.com/zecraid/ShifuDemo.git

# 构建应用
RUN go build -o main ./ShifuDemo/src/task.go

# 启动应用
CMD ["./main"]
