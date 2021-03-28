package base

import (
	"fmt"
	"time"
)

//日期转化练习
func DateFormate() {
	var ms int64 = 1616515200000
	fmt.Println("ms:", ms)
	//2021-03-28 23:06:57.306795 +0800 CST m=+5.977415498
	fmt.Println(time.Now())
	//1616944029 秒级
	fmt.Println(time.Now().Unix())
	//2021-03-28 23:09:50.906652 +0800 CST
	fmt.Println(time.Now().Local())
	//2021-03-28 23:07:09
	fmt.Println(time.Unix(1616944029, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(ms/1000, 0).Format("2006-01-02 15:04:05"))
}
