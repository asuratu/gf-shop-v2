package time

import (
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}
