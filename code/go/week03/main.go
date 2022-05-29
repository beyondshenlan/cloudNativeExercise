package main

import (
	"context"
	"exercise1/ch9/test_ex1/web"
	"fmt"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "这是订单")
}

func main() {
	shutdown := web.NewGracefulShutdown()
	server := web.NewSdkHttpServer("my-test-server",
		web.MetricFilterBuilder, shutdown.ShutdownFilterBuilder)
	adminServer := web.NewSdkHttpServer("admin-test-server",
		web.MetricFilterBuilder, shutdown.ShutdownFilterBuilder)

	// 注册路由
	_ = server.Route("POST", "/user/create/*", web.SignUp)
	_ = server.Route("POST", "/slowService", web.SlowService)

	// 准备静态路由

	//staticHandler := web.NewStaticResourceHandler(
	//	"demo/static", "/static",
	//	web.WithMoreExtension(map[string]string{
	//		"mp3": "audio/mp3",
	//	}), web.WithFileCache(1 << 20, 100))
	//// 访问 Get http://localhost:8080/static/forest.png
	//server.Route("GET", "/static/*", staticHandler.ServeStaticResource)

	go func() {
		if err := adminServer.Start(":8081"); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := server.Start(":8080"); err != nil {
			// 快速失败，因为服务器都没启动成功，啥也做不了
			panic(err)
		}
		// 假设我们后面还有很多动作
	}()
	web.WaitForShutdown(
		func(ctx context.Context) error {
			// 假设我们这里有一个 hook
			// 可以通知网关我们要下线了
			fmt.Println("mock notify gateway")
			time.Sleep(time.Second * 2)
			return nil
		},
		shutdown.RejectNewRequestAndWaiting,
		// 全部请求处理完了我们就可以关闭 server了
		web.BuildCloseServerHook(server, adminServer),
		func(ctx context.Context) error {
			// 假设这里我要清理一些执行过程中生成的临时资源
			fmt.Println("mock release resources")
			time.Sleep(time.Second * 2)
			return nil
		})


}




