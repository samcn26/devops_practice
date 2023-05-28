package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

//模块二作业
//编写一个 HTTP 服务器
//
//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
//当访问 localhost/healthz 时，应返回 200
func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		for name, headers := range r.Header {
			for _, h := range headers {
				w.Header().Add(name, h)
			}
		}

		// Read VERSION from environment variables and write into response header
		version := os.Getenv("VERSION")
		if version != "" {
			w.Header().Set("Version", version)
		}

		// Write the client IP and HTTP status code to server standard output
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		log.Printf("Client IP: %s, HTTP Status code: %d\n", ip, http.StatusOK)

		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
