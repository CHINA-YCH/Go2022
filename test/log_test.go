package test

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

var lo sync.Mutex
var MP = make(map[string]string, 0)

func TestPush(t *testing.T) {
	t.Log(12344)
}
func TestNil2(t *testing.T) {
	var num *int
	var num2 int
	t.Log("num=", num)
	t.Log("num2=", num2)
}

func TestArr(t *testing.T) {
	var ar = []int{1, 2, 3, 4, 5, 6}
	to := len(ar)
	mi := 3
	i := 0
	c := mi - i
	r := ar[to-c]
	t.Log(r)

}

func TestSortStr(t *testing.T) {
	var mp = make(map[string]map[string]string, 0)

	mp["001"] = map[string]string{
		"2022-07-01 01:01:01": "1",
		"2022-07-01 01:02:01": "2",
		"2022-07-01 01:02:02": "3",
	}
	mp["004"] = map[string]string{
		"2022-07-02 03:01:03": "1",
		"2022-08-02 03:01:03": "1",
		"2022-07-08 03:01:03": "1",
	}
	mp["005"] = map[string]string{"2022-07-02 03:02:01": "1"}
	mp["003"] = map[string]string{"2022-07-01 01:01:03": "1"}
	mp["002"] = map[string]string{"2022-07-01 01:01:02": "1"}
	c := map[string]string{"2022-10-10 10:10:10": "1"}
	mp["090"] = c
	var camIdAndTime = make(map[string]string, 0)
	for k, v := range mp {
		var temp = make([]string, 0)
		for kk := range v {
			temp = append(temp, kk)
		}
		sort.SliceStable(temp, func(i, j int) bool {
			if temp[i] <= temp[j] {
				return true
			}
			return false
		})
		var small = temp[0]
		camIdAndTime[k] = small
	}
	t.Log(camIdAndTime)
}

func TestTime2Str(t *testing.T) {
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05")
	t.Log(format)

}

func TestMapNil(t *testing.T) {
	var m1 = make(map[string]string, 0)
	t.Log(m1)
	if len(m1) == 0 {
		t.Log("nil ")
	} else {
		t.Log("not nil")
	}
	var m2 = make(map[string]*int)
	v := 1
	m2["a"] = &v
	i := m2["b"]
	if i == nil {
		t.Log(" nil")
	} else {
		t.Log(" not  nil")
	}

}

func TestQueue(t *testing.T) {
	linkList := list.New()
	linkList.PushBack("a")
	linkList.PushBack("b")
	linkList.PushBack("c")
	linkList.PushBack("d")
	v := linkList.Front().Value // 取出第一个元素
	t.Log(v)

	for head := linkList.Front(); head != nil; head = head.Next() {
		//var value interface{}
		//if head.Prev() != nil {
		//	value = head.Prev().Value
		//}
		//var next interface{}
		//if head.Next() != nil {
		//	next = head.Next().Value
		//}
		//log.Infof("head value = %v, pre value = %v, next value = %v", head.Value, value, next)
		log.Infof("head value = %v", head.Value)
	}

}

func TestLists(t *testing.T) {
	var infoList []string
	infoList = append(infoList, "a")
	infoList = append(infoList, "b")
	infoList = append(infoList, "c")
	delIndex := 1
	infoList = append(infoList[:delIndex], infoList[delIndex+1:]...)
	t.Log(infoList)
	var info2 []string
	info2 = append(info2, "d")
	info2 = append(info2, "e")

	infoList = append(infoList, info2...)

	t.Log(infoList)

}

func TestHttpPost(t *testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 100,
		},
		Timeout: 3 * time.Second, // 超时加在这里，是每次调用的超时(超时时间不应该设置在连接上，而应该设置在请求这个动作上)
	}
	//Post请求示例
	URL := "http://0.0.0.0:8083/v1/dangerous_vehicle/upsert"
	// 表单数据 contentType := "application/x-www-form-urlencoded" json data := "vehiclePlate=浙F3TC08&hazardType="102204""
	contentType := "application/json"
	data := `{"vehiclePlate":"浙F3TC08","hazardType":"102204"}`
	resp, err := client.Post(URL, contentType, strings.NewReader(data))
	if err != nil {
		log.Info("post failed, err:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("get resp failed, err", err)
		return
	}
	log.Info(string(b))
}

func TestHttpGetParam(t *testing.T) {
	// 带参数的GET请求示例
	apiUrl := "http://0.0.0.0:8083/v1/dangerous_vehicle"
	// URL param
	data := url.Values{}
	data.Set("vehiclePlate", "浙F3TC08")
	// 把string转换为 url
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		log.Info("parse url requestUrl failed,err：", err)
	}
	//URL encode
	u.RawQuery = data.Encode()
	log.Info(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		log.Info("get failed, err:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)
	//读取内容
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("get resp failed, err:", err)
		return
	}
	log.Info(string(b))

}

func TestHttpGet(t *testing.T) {
	//发送get请求
	resp, err := http.Get("http://0.0.0.0:8083/v1/dangerous_vehicle")
	if err != nil {
		log.Error(err)
	}
	//关闭Body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)
	//读取body内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("read from resp.Body failed, err", err)
		return
	}
	//输出字符串内容
	log.Infof(string(body))
}

type speedInfo struct {
	Speed *float64 `json:"speed"` // 给 * 默认值不会为""或0 而是nil
	Plate *string  `json:"plate"`
	Lane  *int     `json:"lane"`
}

func (s speedInfo) IsEmpty() bool {
	return reflect.DeepEqual(s, speedInfo{})
}

func TestNil(t *testing.T) {
	var s1 = speedInfo{} // 指针默认不为空 {Speed:0 Plate: Lane:0}
	if s1.Speed == nil {
		t.Logf("- - - - - - - - - -%+v", s1)
	} else {
		t.Logf("===================%+v", s1)
	}

	var s2 = speedInfo{
		Speed: toF(1.0),
	}
	if s2 == (speedInfo{}) {
		log.Infof("s == speedInfo{} empty ")
	}

	if s2.IsEmpty() {
		log.Infof("reflect deep is empty")
	}
	log.Infof("%+v", *s2.Speed)
}

func toF(v float64) *float64 {
	return &v
}

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
