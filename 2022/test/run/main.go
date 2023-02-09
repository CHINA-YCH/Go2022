package main

import (
	"git.supremind.info/gobase/2022/log-d"
	log "github.com/sirupsen/logrus"
)

func main() {
	log_d.SetLog()
	Demo()
}

func Demo() {
	log.Infof("I'm demo %s ...", "^_^")
	log.Debugf("I'm demo %s ...", "^_^")
	log.Warnf("I'm demo %s ...", "^_^")
}
