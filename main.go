package main

import (
	_ "shop/boot/time"
	"shop/internal/cmd"
	_ "shop/internal/logic"
	_ "shop/internal/packed"
	_ "shop/utility/response"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
	cmd.Mq.Run(gctx.GetInitCtx())
}
