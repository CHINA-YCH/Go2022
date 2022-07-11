package test

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var lo sync.Mutex
var MP = make(map[string]string, 0)

func TestLock(t *testing.T) {

	go func() {
		t.Log("- - - - -")
		lo.Lock()
		MP["a"] = "b"
		lo.Unlock()
	}()

	go func() {
		t.Log("xxxxxx")
		lo.Lock()
		MP["a"] = "a"
		lo.Unlock()
	}()

	time.Sleep(3 * time.Second)
	t.Log(MP)
}
func TestDel(t *testing.T) {
	var push = []string{"a", "b", "c", "d"}
	delIndex := 1
	push = append(push[:delIndex], push[(delIndex+1):]...)
	t.Log(push)
}

func TestSort2(t *testing.T) {
	var l = []float64{1.0, 2.0, 0.5, 0.9, 14}
	sort.SliceStable(l, func(i, j int) bool {
		if i >= j {
			return true
		} else {
			return false
		}
	})
	sort.Float64s(l)
	t.Log(l)
}
func TestTime(t *testing.T) {
	now := time.Now().Unix()
	time.Sleep(1 * time.Second)
	now2 := time.Now().Unix()

	dif := now2 - now
	t.Log(dif)
}
func TestPrint(t *testing.T) {
	type PlateInfo struct {
		Speed  float64 `json:"speed" yaml:"speed"`
		Number int     `json:"number" yaml:"number"`
	}
	var tempStruct = &PlateInfo{}
	tempStruct.Speed = 11.2
	tempStruct.Number = 12
	marshal, _ := json.Marshal(tempStruct)
	log.Info(string(marshal))
	str := "2022-06-06 12:09:01 8888"
	s := str[0:19]
	log.Println(s)

	ss := `{"speed":11.2,"number":12}`
	tt := PlateInfo{}
	err := json.Unmarshal([]byte(ss), &tt)
	if err != nil {
		log.Errorf("---%v,orgdata= %v\n", err, ss)
	}

	var a int = 12
	var size int = 10
	var sss int = a / size
	log.Infof("----------%v", sss)
}
func TestFor(t *testing.T) {
	mp := make(map[string]string, 0)
	mp["a"] = "a"
	mp["b"] = "b"
	mp["c"] = "c"
	mp2 := make(map[string]string, 0)
	mp2["a"] = "a"
	for k := range mp {
		if mp2[k] != "" {
			continue
		}
		t.Log("key=", k)
	}
}

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
