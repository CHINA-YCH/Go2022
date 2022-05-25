package main

import log "github.com/sirupsen/logrus"

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceQuote:      true, //键值对加引号
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})

	Demo()
}

func Demo() {
	log.Infof("I'm demo %s ...", "^_^")
	log.Debugf("I'm demo %s ...", "^_^")
	log.Warnf("I'm demo %s ...", "^_^")
}
