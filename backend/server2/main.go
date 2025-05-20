// 反向代理与负载均衡教学示例：后端服务2
// 本服务监听8082端口，接收POST请求并返回带有服务器标识的响应
// 相关知识点：Go http服务、JSON处理、基础API设计
//
// 教学要点：
// 1. 如何用Go实现一个简单的HTTP服务
// 2. 如何解析JSON请求体并返回JSON响应
// 3. 如何区分不同后端服务的响应
//
// 练习：
// - 尝试修改响应内容，前端如何区分不同服务？
// - 增加自定义字段，体验前后端联动

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// 定义前端发送的消息结构体
type Message struct {
    Message string `json:"message"`
}

// 定义后端返回的响应结构体
type Response struct {
    Response string `json:"response"`
}

func main() {
    // 注册根路径的处理函数
    http.HandleFunc("/", handleRequest)
    
    port := 8082
    fmt.Printf("Server 2 starting on port %d...\n", port)
    // 启动HTTP服务
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        log.Fatal(err)
    }
}

// 处理POST请求，返回带有服务器标识的响应
func handleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }
    // 解析JSON请求体
    var msg Message
    if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // 构造响应，带有服务器标识
    response := Response{
        Response: fmt.Sprintf("Server 8082 收到消息: %s", msg.Message),
    }
    // 返回JSON响应
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}