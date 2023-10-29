# go-ip2region-web

使用go查询ip2region.xdb，提供web接口

数据来源和源码参考： https://github.com/lionsoul2014/ip2region

## 运行

```shell
# 源码运行
go run main.go

# docker运行
docker run -it --rm -p 127.0.0.1:8899:8899 --name go-ip2region-web fa1seut0pia/go-ip2region-web
```

## 使用

```shell

# 默认获取请求方的ip地址
curl http://127.0.0.1:8899
# 指定ip地址
curl http://127.0.0.1:8899/223.5.5.5

```

返回json格式，如

```json
{
    "ip": "223.5.5.5",
    "country": "中国",
    "region": "0",
    "province": "浙江省",
    "city": "杭州市",
    "isp": "阿里云"
}
```

## Docker

```shell

# 运行
docker run -it --rm -p 127.0.0.1:8899:8899 --name go-ip2region-web fa1seut0pia/go-ip2region-web

# 有需要时可本地构建
docker build -t go-ip2region-web .

```