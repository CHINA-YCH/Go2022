package main

import (
	logd "git.supremind.info/gobase/log-d"
	log "github.com/sirupsen/logrus"
)

func main() {
	logd.SetLog()
	Demo()
}

func Demo() {
	log.Infof("I'm demo %s ...", "^_^")
	log.Debugf("I'm demo %s ...", "^_^")
	log.Warnf("I'm demo %s ...", "^_^")
}
