package main

import log "github.com/sirupsen/logrus"

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	initLog()
	log.Info("我是一条日志")
	log.WithFields(log.Fields{"key": "value"}).Info("我要打印了")
}
