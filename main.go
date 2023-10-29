package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net/http"
	"strings"
)

func main() {

	dbPath := "ip2region.xdb"
	// 1、从 dbPath 加载整个 xdb 到内存
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return
	}

	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		fmt.Printf("failed to create searcher with content: %s\n", err)
		return
	}

	println("使用方式：")
	println("curl http://127.0.0.1:8899")
	println("curl http://127.0.0.1:8899/223.5.5.5")
	println("curl http://127.0.0.1:8899/223.5.5.5.json")

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		ip := context.ClientIP()
		reg, err := searcher.SearchByStr(ip)
		if err != nil {
			context.String(500, "ip error")
		}
		context.String(200, reg)
	})
	r.GET("/:ip", func(context *gin.Context) {

		ip := context.Param("ip")

		if len(ip) > 5 && ip[len(ip)-5:] == ".json" {
			ip = ip[:len(ip)-5] // 去掉末尾的 ".json"
			reg, err := searcher.SearchByStr(ip)
			if err != nil {
				context.String(500, "ip error")
			}

			ipRegion := strings.Split(reg, "|")
			context.JSON(http.StatusOK, gin.H{
				"ip":       ip,
				"country":  ipRegion[0],
				"region":   ipRegion[1],
				"province": ipRegion[2],
				"city":     ipRegion[3],
				"isp":      ipRegion[4],
			})

		} else {
			reg, err := searcher.SearchByStr(ip)
			if err != nil {
				context.String(500, "ip error")
			}
			context.String(200, reg)
		}

	})
	rErr := r.Run(":8899")
	if rErr != nil {
		return
	}
}
