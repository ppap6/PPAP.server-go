FROM golang:1.13-alpine

# 使用阿里云加速
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates

EXPOSE 8080

RUN go env -w GO111MODULE="auto" \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app

ADD . .

RUN go build -o main main.go

CMD ["./main"]