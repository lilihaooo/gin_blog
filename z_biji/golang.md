## golang
## map
    // 获取map中的值, 返回两次参数 值, 是否存在(bool)
    if conn, exists := ws[re.UserID]; exists {
        err = conn.WriteMessage(websocket.TextMessage, byteSlice) // 文本类型
        if err != nil {
            global.Logrus.Error("Write消息失败:", err)
            break
        }
    }