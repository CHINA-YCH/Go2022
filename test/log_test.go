package test

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	preCarsLane := make(map[string]int, 0)
	if preCarsLane["11"] == 0 {
		t.Log(preCarsLane["11"])
	} else {
		t.Logf("%v----", preCarsLane["11"])
	}
}

func TestLogOutput(t *testing.T) {
	log.Printf("")
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		ForceQuote:      true,                      //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05.000", //时间格式
		FullTimestamp:   true,
	})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.Trace("trace msg")
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("error msg")
	//log-d.Fatal("fatal msg")
	//log-d.Panic("panic msg")

	log.Infof("哈哈哈: %s", "world")
	log.Warnf("警告:%s", "hello")
}

func TestFloat(t *testing.T) {
	var fl = 94075.0
	str := fmt.Sprintf("%.0f", fl)
	log.Infof("str: %s", str)
	after := string([]byte(str)[len(str)-3:])
	log.Infof("after: %s", after)
	pre := strings.Replace(str, after, "", -1)
	log.Infof("pre: %s", pre)
}

func TestSort(t *testing.T) {
	var arr = []int{11, 2, 7, 4, 5, 99}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	log.Infof("arr: %v", arr)
}

func TestCopy(t *testing.T) {
	var arr1 = []int{1, 2, 3}
	var arr2 = []int{11, 22, 33, 44}
	copy(arr1, arr2)
	log.Infof("arr1: %v", arr1)
}

func TestConv(t *testing.T) {
	uid := strconv.FormatInt(int64(111), 10)
	log.Infof("uid: %v", uid)
	cameraId := "cameraId"
	key := fmt.Sprintf(cameraId, "_", 111)
	s := cameraId + "_" + uid
	log.Infof("key: %v", key)
	log.Infof("s: %v", s)
}

func TestRand(t *testing.T) {
	var sed int64 = 100
	var add int64 = 0
	rand.Seed(time.Now().UnixNano())
	i := rand.Int63n(sed) + add
	log.Infof("i=%v", i)
}

func TestFormat(t *testing.T) {
	var uid int = 12
	var cameraId = "1234"
	key := fmt.Sprintf("%s_%v", cameraId, uid)
	log.Infof("key: %v", key)
}
