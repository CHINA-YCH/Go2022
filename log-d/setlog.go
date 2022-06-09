package log_d

import log "github.com/sirupsen/logrus"

func SetLog() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceQuote:      true, //键值对加引号
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})
}
