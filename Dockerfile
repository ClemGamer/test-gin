# syntax=docker/dockerfile:1
FROM golang:1.20.3-alpine
WORKDIR /src
# 依照package manager路徑設定快取,只有第一次build需要下載所有相依的package
# 在沒清除快取的狀況重新build,只會下載有更新的package
# 清除快取：docker builder prune -af
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sun \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
COPY . .
RUN go build -o /bin/test-server .

# scratch是docker最小的base image
FROM scratch
# 只把第0層build好的binary複製到scratch
COPY --from=0 /bin/test-server /bin/
ENTRYPOINT ["/bin/test-server"]

# 最後build出來的是含有binary的scratch
# build階段都不會保留