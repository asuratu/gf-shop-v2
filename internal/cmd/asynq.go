package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Mq = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("start mq")
			return nil
		},
	}
)
