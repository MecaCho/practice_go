

# nil channel

A1. 对一个 nil channel执行发送操作会一直阻塞。

A2. 对一个 nil channel执行接受操作会一直阻塞。

A3. 发送到关闭的channel会引起panic

A4. 从关闭的cannel读操作，会立刻返回数据0值。