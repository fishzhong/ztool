package service

import (
	"fmt"
	sdk "github.com/sensorsdata/sa-sdk-go"
	"time"
)

const Url = "http://sensorsdata.web.bigdata.fulu.com:18106/sa?project=yunshi_pro"

func Send() {
	// 从神策分析配置页面中获取数据接收的 URL
	//SA_SERVER_URL := "YOUR_SERVER_URL"
	// 初始化一个 Consumer，用于数据发送
	// DefaultConsumer 是同步发送数据，因此不要在任何线上的服务中使用此 Consumer
	consumer, _ := sdk.InitDebugConsumer(Url, true, 10000)
	//consumer2,_:=sdk.InitConcurrentLoggingConsumer()
	//...
	// 使用 Consumer 来构造 SensorsAnalytics 对象
	sa := sdk.InitSensorsAnalytics(consumer, "yunshi_pro", false)
	defer sa.Close()

	//distinct_id := "ABCDEF1234567"

	properties := map[string]interface{}{
		// "$time" 属性是系统预置属性，取值为 int64 类型，表示事件发生的时间，如果设置该属性，则替换 SDK 内部获取的系统当前时间。如果不填入该属性，则默认使用系统当前时间
		//"$time": time.Now().Unix(),
		//// "$ip" 属性是系统预置属性，如果服务端中能获取用户 IP 地址，并填入该属性，神策分析会自动根据 IP 地址解析用户的省份、城市信息
		//"$ip": "123.123.123.123",
		//// 商品 ID
		//"ProductId": "123456",
		//// 商品类别

		//"ProductCatalog": "Laptop Computer",
		//// 是否加入收藏夹，Boolean 类型的属性
		//"IsAddedToFav": true,
	}

	properties = map[string]interface{}{
		"userid":      "3223343443",
		"status":      "1",
		"use_seconds": int(12222),
		//"create_date": time.Now().Format("2006-01-02 15:04:05"),
	}

	properties = map[string]interface{}{
		"userid":         "2756794314752",
		"channel_number": "",
		"register_from":  1,
		"register_time":  time.Now(),
		"phone":          "19983344913",
		"custom":         "yebao",
		//"use_seconds": int(12222),
		//"create_date": time.Now().Format("2006-01-02 15:04:05"),
	}
	err := sa.ProfileSet("2756794314752", properties, true)
	//err := sa.Track("32233434431", "open_close_vip_operation", properties, true)

	//err := sa.Track(distinct_id, "GwView", properties, true)

	fmt.Println(err)
}
