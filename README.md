# go-ip2region-web

使用go查询ip2region.xdb，提供web接口

数据来源和源码参考： https://github.com/lionsoul2014/ip2region

## 使用

```shell

# 默认获取请求方的ip地址
curl http://127.0.0.1:8899
# 指定ip地址
curl http://127.0.0.1:8899/223.5.5.5
# 指定ip地址以json的格式返回
curl http://127.0.0.1:8899/223.5.5.5.json

```

## 构建

```shell
# Windows 
go build -o dist/go-ip2region-web.exe

# Linux
# 如果此时是在windows中构建Linux平台二进制文件需执行，否则忽略下方的环境配置
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux

go build -o dist/go-ip2region-web-linux-amd64

# 还原
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=windows
go env -w GOOS=windows
```

## Docker

```shell
# 构建
docker build -t go-ip2region-web .

# 运行
docker run -it -p 127.0.0.1:8899:8899 --name ipr go-ip2region-web
```