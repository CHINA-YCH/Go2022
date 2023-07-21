package s3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/golib/assert"
)

/*
 * @Author: veno
 * @File: s3_test
 * @Version: ...
 * @Date: 2023-04-23 10:21:47
 * @Description: ...
 */

func TestS3Image(t *testing.T) {
	client, err := NewS3FileServer(&FileServerConfig{
		Endpoint: aws.String("10.15.2.22:9000"),
		AK:       aws.String("smaitestak"),
		SK:       aws.String("smaitestsk"),
		Bucket:   aws.String("default"),
		Term:     aws.String("short-term"),
	})
	// 发送 HTTP GET 请求获取图片
	resp, err := http.Get("http://10.15.2.34:9720/1.jpg")
	if err != nil {
		fmt.Println("Failed to send request", err)
		return
	}
	defer resp.Body.Close()
	// 读取请求响应内容为字节数组
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Failed to read response", err2)
		return
	}
	assert.Nil(t, err)
	retUrl, err := client.Save("/test/1.jpg", data) // key 文件名 data 存储的内容
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ret rul: ", retUrl)
	assert.Nil(t, err)
	t.Logf("s3 key for testfile.txt:%v", retUrl)
}

func TestS3Put(t *testing.T) {
	client, err := NewS3FileServer(&FileServerConfig{
		Endpoint: aws.String("10.15.2.22:9000"),
		AK:       aws.String("smaitestak"),
		SK:       aws.String("smaitestsk"),
		Bucket:   aws.String("default"),
	})
	data := []byte("this is a test file")
	if err != nil {
		fmt.Println(err)
	}

	assert.Nil(t, err)
	key, err := client.Save("testfile.txt", data) // key 文件名 data 存储的内容
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
	t.Logf("s3 key for testfile.txt:%v", key)

}

func TestS3Img(t *testing.T) {
	// 打开本地图片文件 http://10.15.2.34:9720/1.jpg
	file, err := os.Open("image.jpg")
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}
	defer file.Close()

	// 读取文件内容为字节数组
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	fmt.Printf("File size: %d bytes\n", len(bytes))
	fmt.Println("- - - - - - - - - - - ")

	// 发送 HTTP GET 请求获取图片
	resp, err := http.Get("https://example.com/image.jpg")
	if err != nil {
		fmt.Println("Failed to send request", err)
		return
	}
	defer resp.Body.Close()

	// 读取请求响应内容为字节数组
	bytes2, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Failed to read response", err2)
		return
	}

	fmt.Printf("File size: %d bytes\n", len(bytes2))
}
