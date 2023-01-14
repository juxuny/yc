# builder
FROM golang:1.16.4 as builder
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src && go env -w GOPROXY=https://goproxy.cn && go mod download
COPY . /src/
RUN cd /src/services/cos/server && CGO_ENABLED=0 go build -o app

# binary
FROM alpine:3.8
RUN apk add --no-cache ca-certificates tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
WORKDIR /app
COPY --from=builder /src/services/cos/server/app /app/entrypoint
ENTRYPOINT /app/entrypoint serve
