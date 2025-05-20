// 反向代理与负载均衡教学示例：代理服务器
// 本服务监听3000端口，负责请求转发和静态文件服务
// 相关知识点：Go反向代理、路由分发、静态资源服务
//
// 教学要点：
// 1. 如何用Go实现反向代理（httputil.NewSingleHostReverseProxy）
// 2. 如何根据路径分发请求到不同后端
// 3. 如何提供静态文件服务
// 4. 如何将根路径重定向到前端页面
//
// 练习：
// - 增加新的后端服务，扩展代理路由
// - 实现简单的负载均衡策略（如轮询）

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

    // 设置路由，将不同API请求转发到不同后端
    http.HandleFunc("/api/server8081", server1Handler)
    http.HandleFunc("/api/server8082", server2Handler)
    
    // 提供静态文件服务（如前端页面、样式、脚本等）
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // 根路径重定向到 index.html
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

// 创建反向代理处理函数，将请求转发到目标后端
func createProxy(target string) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // 解析目标URL
        targetUrl, _ := url.Parse(target)
        
        // 创建反向代理
        proxy := httputil.NewSingleHostReverseProxy(targetUrl)
        
        // 设置请求头，转发请求
        r.URL.Host = targetUrl.Host
        r.URL.Scheme = targetUrl.Scheme
        r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
        
        // 代理处理请求
        proxy.ServeHTTP(w, r)
    }
}