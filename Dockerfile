FROM golang:1.21 as builder

WORKDIR /src

COPY . .

ENV GOPROXY=https://goproxy.cn,direct

RUN rm -rf dev.go go.work go.work.sum ; go mod tidy
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /src/door


FROM alpine:latest as compress

WORKDIR /

COPY --from=builder /src/door /door

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories 
RUN  apk add --no-cache upx ca-certificates tzdata \
  && upx -5 door -o _upx_door \
  && mv -f _upx_door door

FROM busybox:stable-glibc

COPY --from=compress /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=compress /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=compress /door /app/door
COPY ./config.prod.toml /app/door.toml

ENTRYPOINT [ "/app/door" ]