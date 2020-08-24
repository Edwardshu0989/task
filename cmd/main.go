package main

import (
	Alog "awesomeProject/Log"
	"awesomeProject/config"
	"awesomeProject/dl"
	"fmt"
	"github.com/prometheus/common/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	Alog.Info("log init")
	fmt.Println(config.Env.ListenAddr)
	_, closeFunc, err := dl.InitApp()
	if err != nil {
		panic(fmt.Errorf("init err %s", err.Error()))
	}
	fmt.Println("server start")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("klt-center exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
