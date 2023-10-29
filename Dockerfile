FROM alpine:latest
LABEL authors="falseutopia#163.com"

WORKDIR /app

COPY dist/go-ip2region-web-linux-amd64 /app/ip2region-web
COPY ip2region.xdb /app/ip2region.xdb

EXPOSE 8899

CMD ["/app/ip2region-web"]