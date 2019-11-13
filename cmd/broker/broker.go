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
	conf, err := config.New("config.yaml")
	if err != nil {
		log.Println("new config failed, err msg is", err)
		return
	}

	lg := logger.Init(conf.Common.LogLevel)
	lg.Info("config", zap.Any("conf", conf))

	s := broker.New(conf, lg)
	if err := s.Start(); err != nil {
		lg.Error("broker start failed", zap.Error(err))
		return
	}

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)

	lg.Info("broker received Stop signal", zap.Any("signal", <-chSig))

	s.Close()
	return
}
