package exit

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitElegantExit(c chan os.Signal) {
	done := make(chan bool, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// 接收到信号返回
		<-c
		done <- true
	}()
	<-done
}
