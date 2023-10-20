FROM golang:alpine as builder
COPY . /usr/local/go/src/github-release-getter
WORKDIR /usr/local/go/src/github-release-getter
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go build -o /usr/bin/github-release-getter github-release-getter

###
FROM alpine as final
ENTRYPOINT ["/usr/bin/github-release-getter"]
COPY --from=builder /usr/bin/github-release-getter /usr/bin/