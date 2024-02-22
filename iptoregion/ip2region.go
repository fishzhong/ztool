package iptoregion

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"strings"
	"time"
)

func ip(ip string) {
	var dbPath = "ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}

	defer searcher.Close()

	// do the search

	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}
	//xdb 支持亿级别的 IP 数据段行数，默认的 region 信息都固定了格式：国家|区域|省份|城市|ISP，缺省的地域信息默认是0。
	//region 信息支持完全自定义，例如：你可以在 region 中追加特定业务需求的数据，例如：GPS信息/国际统一地域信息编码/邮编等。
	//也就是你完全可以使用 ip2region 来管理你自己的 IP 定位数据。

	fmt.Printf("{region: %s, took: %d}\n", region, time.Since(tStart).Nanoseconds())
	ts := strings.Split(region, "|")
	if len(ts) != 5 {
		fmt.Println("IP未知")
	} else {
		var s string
		for i, t := range ts {
			if t == "中国" || i == 4 {
				continue
			}
			if t != "0" {
				s = s + t
			}
		}
		fmt.Println("位置:", s)
	}
}
