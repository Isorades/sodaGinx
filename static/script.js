// 反向代理与负载均衡教学示例：前端脚本
// 本脚本负责与代理服务器通信，发送消息并动态展示响应历史
// 相关知识点：fetch API、DOM 操作、异步编程
//
// 教学要点：
// 1. 如何用 fetch 发送 POST 请求到代理服务器
// 2. 如何处理 JSON 响应并更新页面内容
// 3. 如何动态维护通信历史
//
// 练习：
// - 增加错误处理，体验异常场景
// - 扩展历史记录展示内容

// 用于动态更新历史记录
function updateHistory(result, server) {
    const historyList = document.getElementById("history");
    const newItem = document.createElement("li");
    newItem.textContent = `Result: ${result}, Processed by: ${server}`;
    historyList.appendChild(newItem);
}

// 添加表单提交处理
document.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector('form');
    form.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const formData = new FormData(form);
        const response = await fetch(`/proxy?${new URLSearchParams(formData)}`);
        const result = await response.text();
        
        // 解析结果并更新历史
        const [resultText, serverText] = result.split(', ');
        const resultNumber = resultText.split(': ')[1];
        const serverName = serverText.split(': ')[1];
        
        updateHistory(resultNumber, serverName);
    });
});

async function sendMessage() {
    const message = document.getElementById('messageInput').value;
    const server = document.getElementById('serverSelect').value;
    
    if (!message.trim()) {
        alert('请输入消息内容！');
        return;
    }

    try {
        // 通过代理API发送POST请求
        const response = await fetch(`/api/server${server}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ message: message }),
        });

        // 解析JSON响应
        const data = await response.json();
        addMessageToHistory(message, data.response, server);
        document.getElementById('messageInput').value = '';
    } catch (error) {
        console.error('Error:', error);
        alert('发送消息失败！');
    }
}

// 动态添加历史记录到页面
function addMessageToHistory(message, response, server) {
    const historyDiv = document.getElementById('messageHistory');
    const timestamp = new Date().toLocaleString();
    
    const messageItem = document.createElement('div');
    messageItem.className = 'message-item';
    messageItem.innerHTML = `
        <div class="timestamp">${timestamp} - 服务器 ${server}</div>
        <div>发送: ${message}</div>
        <div>响应: ${response}</div>
    `;
    
    historyDiv.insertBefore(messageItem, historyDiv.firstChild);
}
