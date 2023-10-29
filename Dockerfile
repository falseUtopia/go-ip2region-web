FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN export GOPROXY=https://proxy.golang.com.cn,direct && go mod download
COPY main.go .
RUN go build -o go-ip2region-web

FROM alpine:latest
LABEL authors="falseutopia#163.com"
WORKDIR /app
COPY --from=build /app/go-ip2region-web .
COPY ip2region.xdb .
EXPOSE 8899
CMD ["/app/go-ip2region-web"]