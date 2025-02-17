package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    // 创建多个后端服务器的代理处理器
    server1Handler := createProxy("http://localhost:8081")
    server2Handler := createProxy("http://localhost:8082")

    // 设置路由
    http.HandleFunc("/api/server8081", server1Handler)
    http.HandleFunc("/api/server8082", server2Handler)
    
    // 提供静态文件服务
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // 将根路径重定向到 index.html
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.ServeFile(w, r, "static/index.html")
            return
        }
        http.NotFound(w, r)
    })

    // 启动代理服务器
    port := 3000
    fmt.Printf("Proxy server starting on port %d...\n", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        log.Fatal(err)
    }
}

func createProxy(target string) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // 解析目标URL
        targetUrl, _ := url.Parse(target)
        
        // 创建反向代理
        proxy := httputil.NewSingleHostReverseProxy(targetUrl)
        
        // 设置请求头
        r.URL.Host = targetUrl.Host
        r.URL.Scheme = targetUrl.Scheme
        r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
        
        // 转发请求
        proxy.ServeHTTP(w, r)
    }
} 