# 🌞 Golang 封装的 aqicn.org API 

`aqcn-go` 实现了 aqicn.org 的 city_feed 查询接口

## 使用方法

a. 注册 Token 

调用 aqicn 接口必须要注册 Token，不过 aqicn 免费提供 Token 注册，并且流程非常简单，
具体请访问 [http://aqicn.org/data-platform/token/#/](http://aqicn.org/data-platform/token/#/)

b. 调用 API

下载 aqicn-go 

```sh
go get github.com/ltyyz/aqicn-go
```

调用示例

```go
package main

import (
	"fmt"
	"github.com/ltyyz/aqicn-go"
)

func main() {
	info, e := aqicn.GetWeatherInfo("YourToken", "shanghai")
	fmt.Println(info, e)
}
```

c. 数据项说明

```go
type WeatherInfo struct {
	Status string  // 返回状态码
	Data   struct {
		Aqi          int // 空气指数
		Idx          int // 城市id
		Attributions []struct { // EPA 信息
			URL  string
			Name string
		}
		City struct {
			Geo  []float64 // 经纬度
			Name string    // 城市名
			URL  string    // 网址
		}
		Dominentpol string //
		Iaqi        struct {
			Co struct {
				V float64 // CO
			}
			H struct {
				V float64 // 湿度
			}
			O3 struct{
				V float64 // O3
			}
			No2 struct {
				V float64 // NO2
			}
			P struct {
				V float64 // 空气压力
			}
			Pm10 struct {
				V float64 // PM10
			}
			Pm25 struct {
				V float64 // PM2.5
			}
			So2 struct {
				V float64 // SO2
			}
			T struct {
				V float64 //  温度
			}
			W struct {
				V float64 // 风力
			}
		}
		Time struct {
			S  string // 更新时间
			Tz string // 时区
			V  int    // 时间值
		}
		Debug struct {
			Sync time.Time // 
		}
	}
}
```

## 感谢

- 动态 json 处理 [gjson](https://github.com/tidwall/gjson)
