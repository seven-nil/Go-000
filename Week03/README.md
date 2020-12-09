# Week3 作业

**基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。**

## 实现

```go
package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var (
		eg         = errgroup.Group{}
		serverErr  = make(chan error, 1)
		httpServer = http.Server{Addr: "0.0.0.0:8080"}
		signalChan = make(chan os.Signal, 1)
		egErr      error
	)

	eg.Go(func() error {
		go func() {
			serverErr <- httpServer.ListenAndServe()
		}()
		select {
		case err := <-serverErr:
			close(signalChan)
			close(serverErr)
			return err
		}
	})

	eg.Go(func() error {
		signal.Notify(signalChan, syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
		<-signalChan
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	egErr = eg.Wait()
	if egErr != nil {
		log.Fatal(egErr)
	}

}

```
