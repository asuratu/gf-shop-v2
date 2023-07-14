package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/os/gcmd"

	"shop/boot/job"
)

var (
	Mq = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("start mq")

			// 初始化 AsynqServer
			job.AsynqServer = job.NewAsynqServer(ctx)

			cronJob := job.NewCronJob(ctx)
			mux := cronJob.Register()
			if err := job.AsynqServer.Run(mux); err != nil {
				os.Exit(1)
			}

			return nil
		},
	}
)
