package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

var VERSION string

func main() {

	os.Setenv("VERSION", "exercise2")
	VERSION = os.Getenv("VERSION")
	if VERSION != "" {
		fmt.Println("my print VERSION" + VERSION)
	} else {
		fmt.Println("no found version")
	}

	//http.HandleFunc("/", responseHander)
	//
	//http.HandleFunc("/healthz", healthz)
	//
	//http.ListenAndServe("127.0.0.1:8888", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", responseHander)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}

}

func responseHander(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("HEAD:", r.Header)
	//2_1 接收客户端 request，并将 request 中带的 header 写入 response header
	var responseHeader http.Header = w.Header()
	for headkey, headvalue := range r.Header {
		var headvalueStr string
		if headvalue != nil {
			headvalueStr = strings.Join(headvalue, ",")
		}
		responseHeader.Add(headkey, headvalueStr)
	}
	//2_2 VERSION 配置，并写入 response header
	responseHeader.Add("Version", VERSION)

	//2_3
	// ip地址
	ipAddr := ClientIP(r)
	//返回码
	//sCode := r.Response.StatusCode
	fmt.Printf("本次访问IP地址是:%s ", ipAddr)
	//fmt.Printf("本次访问IP地址是:%s , 返回状态码是:%d", ipAddr, sCode)

	w.Write([]byte("<html><center> <font size=\"40\">Hello！模块二测试</font></center></html>"))
	//fmt.Println("A client has just visited!")
}

//2_4
func healthz(w1 http.ResponseWriter, r *http.Request) {
	w1.Write([]byte("200"))
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
//解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
