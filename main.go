package main

import (
	"os"
	"time"
	"syscall"
	"context"
	"net/http"
	"os/signal"

	"github.com/liangran2018/lived/system"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

func main()  {
	app := gin.Default()

	// 设置panic恢复
	app.Use(system.AddRecovery())

	// 添加路由
	system.AddRoute(app)

	// 绑定server
	server := &http.Server{
		Addr:           ":8888",
		Handler:        app,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 异步运行server
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("启动程序失败: " + err.Error())
		}
	}()

	//新建日志文件
	log.NewLogFile()
	log.GetLogger().Log(log.Info, "Program Start!!!")
	// 监听系统信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// 定时任务
	for {
		select {
		case <-ch:
			log.GetLogger().Log(log.Info, "Program Exit!!!")
			server.Shutdown(context.TODO())
			signal.Stop(ch)
			os.Exit(0)
		}
	}
}


