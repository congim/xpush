package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/congim/xpush/broker"
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/provider/logger"
	"go.uber.org/zap"
)

func main() {
	var conf *config.Config
	var err error
	if len(os.Args) >= 2 {
		conf, err = config.New(os.Args[1])
	} else {
		conf, err = config.New("config.yaml")
	}
	if err != nil {
		log.Println("new config failed, err msg is", err)
		return
	}

	brokerLog := logger.Init(conf.Common.LogLevel)
	brokerLog.Info("config", zap.Any("conf", conf))

	brokerServer := broker.New(conf, brokerLog)
	if err := brokerServer.Start(); err != nil {
		brokerLog.Error("broker start failed", zap.Error(err))
		return
	}

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)

	brokerLog.Info("broker received Stop signal", zap.Any("signal", <-chSig))

	brokerServer.Close()
	return
}
