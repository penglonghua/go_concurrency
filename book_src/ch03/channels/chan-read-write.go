package main

// chan 中有 读，也有写,写就是生产者，读就是消费者.
// 问题，能否在 生产者中进行消费，消费者中进行生产那?
// 答案是不能, 看下 报错信息

// 报错信息为:
//# command-line-arguments
//./chan-read-write.go:18:2: invalid operation: <-writeStream (receive from send-only type chan<- interface {})
//./chan-read-write.go:19:13: invalid operation: readStream <- struct {} literal (send to receive-only type <-chan interface {})

func main() {

	// 生产者
	writeStream := make(chan <- interface{}) //生产者
	readStream := make(<- chan interface{} ) //消费者

	<-writeStream
	readStream <- struct {}{}


}
