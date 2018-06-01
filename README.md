# Lunar
Go语言实现的阳历阴历转换

# 安装
```
go get github.com/superwen/lunar
```

# 使用
以下是代码示例：
```
package main

import (
	"fmt"
	"time"
	"github.com/superwen/lunar"
)

func main()  {
	myLunar1 , err := lunar.SolarTimeToLunar(time.Now().AddDate(0,1,0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", myLunar1)

	myLunar2 , err := lunar.SolarToLunar(2018, 5,20)
    if err != nil {
    	fmt.Println(err.Error())
    	return
    }
    fmt.Printf("%#v\n", myLunar2)
}
```
说明：
- 本阴历支持的最小年份为 1891，最大年份为 2100
- Lunar 结构的属性有
  Year 年份
  Month 月份（中文）
  Date 日（中文）
  Nian 天干地支
  EMonth 月份（数字）
  EDate 日（数字）
  Zodiac 生效

