package main

import (
	"fmt"
	"github.com/spf13/viper"
	"web-app/controllers"

	//"web-app/dao/mysql"
	"web-app/pkg/snowflake"

	//"github.com/spf13/viper"
	"context"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-app/Logger"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/routes"
	"web-app/settings"
)

//go常用开发脚手架模板
func main() {
	//加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Println("init setting faild err", err)
		return
	}
	fmt.Println("init setting success")
	//初始化日志
	if err := Logger.Init(viper.GetString("app.mode")); err != nil {
		fmt.Println("init Logger faild err", err)
		//return
	}
	defer zap.L().Sync()
	fmt.Println("init logger success")
	//zap.L().Debug("logger init success")
	//初始化mysql

	err := mysql.InitDB()
	if err != nil {
		zap.L().Error("init mysql faild err:", zap.Error(err))
		fmt.Println("数据库连接失败", err)
		return
	}
	defer mysql.Close()
	fmt.Println("init mysql success")
	//初始化redis连接
	err = redis.Init()
	if err != nil {
		fmt.Println("init redis faild err", err)
		return
	}
	defer redis.Close()
	fmt.Println("init redis success")
	if err := snowflake.Init(viper.GetString("app.starttime"), viper.GetInt64("app.machineid")); err != nil {
		fmt.Println("init snowflake faild ", err)
		return
	}

	////注册路由
	err = controllers.InitTrans("zh")
	if err != nil {
		fmt.Println("controllers.InitTrans err:", err)
		return

	}
	r := routes.SetupRoute()
	srv := &http.Server{
		Addr:    viper.GetString("app.port"),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Info("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
	//启动服务（优雅关机和）
}
