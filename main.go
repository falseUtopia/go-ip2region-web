package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net/http"
	"strings"
)

type IPRegion struct {
	IP       string `json:"ip"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Province string `json:"province"`
	City     string `json:"city"`
	ISP      string `json:"isp"`
}

func regionStr2Region(ip string, regionStr string) IPRegion {

	ipRegion := strings.Split(regionStr, "|")

	region := IPRegion{
		IP:       ip,
		Country:  ipRegion[0],
		Region:   ipRegion[1],
		Province: ipRegion[2],
		City:     ipRegion[3],
		ISP:      ipRegion[4],
	}

	return region

}

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

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {

		ip := context.ClientIP()
		reg, err := searcher.SearchByStr(ip)
		if err != nil {
			context.String(500, "ip error")
			return

		}

		context.JSON(http.StatusOK, regionStr2Region(ip, reg))

	})
	r.GET("/:ip", func(context *gin.Context) {

		ip := context.Param("ip")
		reg, err := searcher.SearchByStr(ip)
		if err != nil {
			context.String(500, "ip error")
			return

		}

		context.JSON(http.StatusOK, regionStr2Region(ip, reg))

	})
	rErr := r.Run(":8899")
	if rErr != nil {
		return
	}
}
