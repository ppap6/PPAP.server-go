FROM golang:1.13-alpine as builder

# 使用阿里云加速
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates

RUN go env -w GO111MODULE="auto" \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app

ADD . .

RUN go build -mod vendor -o /app/main .


FROM alpine:latest as final

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]