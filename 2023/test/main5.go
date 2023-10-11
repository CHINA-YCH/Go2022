package main

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

/*
 * @Author: veno
 * @File: main5
 * @Version: ...
 * @Date: 2023-07-26 15:30:10
 * @Description: ...
 */

func main() {
	TimeFormat := "2006-01-02 15:04:05"
	now := time.Now()
	s := now.Format(TimeFormat)
	log.Info("-----", s)
	split1 := strings.Split(s, " ")
	dateSplit := strings.Split(split1[0], "-")
	year := dateSplit[0]
	monthDay := dateSplit[1] + dateSplit[2]
	hour := strings.Split(split1[1], ":")[0]

	log.Info("year---", year)
	log.Info("monthDay---", monthDay)
	log.Info("hour---", hour)

	goString := time.Now().UnixMicro()
	goString1 := time.Now().UnixMicro()
	time.Sleep(1 * time.Nanosecond)
	goString2 := time.Now().UnixMicro()
	goString3 := time.Now().UnixMicro()
	log.Info("----", goString)
	log.Info("----", goString1)
	log.Info("----", goString2)
	log.Info("----", goString3)

	tt := "hsdofjs-snfaodnkak.jpg"
	split := strings.Split(tt, "$")
	log.Info(split[0])

}
