// 用于动态更新计算历史
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
        const response = await fetch(`/api/server${server}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ message: message }),
        });

        const data = await response.json();
        addMessageToHistory(message, data.response, server);
        document.getElementById('messageInput').value = '';
    } catch (error) {
        console.error('Error:', error);
        alert('发送消息失败！');
    }
}

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
