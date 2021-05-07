package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// ➂xの値をcに送信する。初期値は0。その後x, y = y, x+yの値を加える。
		case c <- x:
			x, y = y, x+y
		// ⓺quitを受信できる場合はquitを出力して終了する。
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// 非同期通信のイメージ(ajaxのweb側)
	// ➀goroutine開始
	go func() {
		for i := 0; i < 10; i++ {
			// ➃<-cで受信した値を出力する。
			fmt.Println(<-c)
		}
		// ⓹➃が10回繰り返されたらquitに0を送信する。
		quit <- 0
	}()
	// ➁関数呼び出し
	fibonacci(c, quit)
}
