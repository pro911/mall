package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"log"
	"mall/initializ"
	"mall/internal/pkg/validators"
	"mall/logger"
	"mall/routers"
	"mall/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "f", "", "传入配置文件路径")
	//解析命令行参数
	flag.Parse()
	fmt.Printf("configFile:%v\n", configFile)

	//1.加载配置文件
	if err := settings.Init(configFile); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}

	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync() //当程序退出是执行同步
	zap.L().Debug("logger init success...")

	////3.初始化MySQL连接
	//if err := initializ.MySQL(settings.Conf.MySQLConfig); err != nil {
	//	fmt.Printf("init MySQL failed,err:%v\n", err)
	//	return
	//}
	//defer initializ.MySQLClose() //当程序停止时关闭mysql资源

	////4.初始化Redis连接
	//if err := initializ.Redis(settings.Conf.RedisConfig); err != nil {
	//	fmt.Printf("init Redis failed,err:%v\n", err)
	//	return
	//}
	//defer initializ.RedisClose()

	//5.初始化SQLite连接
	if err := initializ.SQLite(settings.Conf.SQLiteConfig); err != nil {
		fmt.Printf("init SQLite failed,err:%v\n", err)
		return
	}
	log.Printf("sssss:%v : %T\n", settings.Conf.AppConfig.JwtSecret, settings.Conf.AppConfig.JwtSecret)

	//初始化gin框架内置的校验器使用的翻译器
	if err := validators.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err%v\n", err)
		return
	}

	//6.注册路由
	r := routers.Init()

	//6.启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.HttpServer.Port),
		Handler: r,
	}
	fmt.Printf("httpServer.addr:%s\n", fmt.Sprintf(":%d", settings.Conf.HttpServer.Port))

	go func() {
		//开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen", zap.Error(err))
		}
	}()

	//等待中断信号来优雅地关闭服务器,为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) //创建一个接收信号的通道
	//kill 默认会发送 syscall.SIGTERM 信号
	//kill -2 发送 syscall.SIGINT 信号,我们常用的Ctrl+C就是触发系统SIGINT信号
	//kill -9 发送 syscall.SIGKILL 信号,但是不能被捕获,所以不需要添加它
	//signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //此处不会阻塞
	<-quit                                               //阻塞在此,当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	//创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//5秒内优雅关闭服务(将未处理完的请求处理完再关闭服务),超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
