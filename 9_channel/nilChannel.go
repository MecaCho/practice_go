package __channel

import "fmt"

////A1. 对一个 nil channel执行发送操作会一直阻塞。
func SendNilChannel() {
	var ch chan int
	ch <- 1
	<-ch
}

////A2. 对一个 nil channel执行接受操作会一直阻塞。
func ReadNilChannel() {
	var ch chan int
	<-ch
}

//A3. 发送到关闭的channel会引起panic
func SendCloseChannel() {

	ch := make(chan int, 1)
	ch <- 0
	fmt.Println("SendCloseChannel", <-ch)

	close(ch)

	ch <- 1

}

//A4. 从关闭的cannel读操作，会立刻返回数据0值。
func ReadCloseChannel() {
	ch := make(chan int, 1)
	ch <- 0
	fmt.Println("ReadCloseChannel", <-ch)

	close(ch)

	fmt.Println(<-ch)

}
