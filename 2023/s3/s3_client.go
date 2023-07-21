package s3

/*
 * @Author: veno
 * @File: s3_client
 * @Version: ...
 * @Date: 2023-04-23 10:19:20
 * @Description: ...
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type IS3ServerClient interface {
	Init() error
	Save(key string, data []byte) (url string, err error)
	SaveKeepForever(key string, data []byte) (url string, err error)
	Get(url string) ([]byte, error)
}

type FileServerConfig struct {
	Region           *string `json:"region,omitempty"`
	Endpoint         *string `json:"endpoint,omitempty"`
	OutEndpoint      *string `json:"out_endpoint,omitempty"`
	DisableSSL       *bool   `json:"disable_ssl,omitempty"`
	S3ForcePathStyle *bool   `json:"s3_force_path_style,omitempty"`
	Bucket           *string `json:"bucket,omitempty"`
	PathPrefix       *string `json:"path_prefix,omitempty"`
	Expire           *int    `json:"expire,omitempty"`
	AK               *string `json:"ak,omitempty"`
	SK               *string `json:"sk,omitempty"`
	RetryTime        *int    `json:"retry_time,omitempty"`
	Term             *string `json:"term,omitempty"`
	PathPrefixFormat *string `json:"path_prefix_format,omitempty"`
}

type FileServer struct {
	baseCfg *FileServerConfig
	config  *aws.Config
	session *session.Session
	service *s3.S3
}

var _DefaultS3FileServerConfig = FileServerConfig{
	Region:           aws.String("cn-east-1"),
	Endpoint:         aws.String("http://localhost:10066"),
	DisableSSL:       aws.Bool(true),
	S3ForcePathStyle: aws.Bool(true),
	Bucket:           aws.String("default"),
	PathPrefix:       aws.String(""),
	Expire:           aws.Int(0),
	RetryTime:        aws.Int(3),
	Term:             aws.String("short-term"),
}

func NewSingleS3Fileserver(config *FileServerConfig) (IS3ServerClient, error) {
	baseCfg := _DefaultS3FileServerConfig
	baseCfg.Merge(*config)
	staticCredentials := credentials.NewStaticCredentials(*(baseCfg.AK), *(baseCfg.SK), "")

	s3Config := aws.Config{
		Region:           baseCfg.Region,
		Endpoint:         baseCfg.Endpoint,
		Credentials:      staticCredentials,
		DisableSSL:       baseCfg.DisableSSL,
		S3ForcePathStyle: baseCfg.S3ForcePathStyle,
	}
	options, err := session.NewSessionWithOptions(session.Options{Config: s3Config})
	if err != nil {
		return nil, err
	}
	return &FileServer{
		baseCfg: &baseCfg,
		config:  &s3Config,
		session: options,
		service: s3.New(options),
	}, nil
}

func (o *FileServer) Init() error {
	_, err := o.service.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(*(o.baseCfg.Bucket)),
	})
	return err
}

func (o *FileServer) Save(filePath string, data []byte) (urlStr string, err error) {
	PathPrefix := ""
	if o.baseCfg.PathPrefixFormat != nil {
		PathPrefix = time.Now().Format(*o.baseCfg.PathPrefixFormat)
	}
	// fullKey 默认为文件名
	fullKey := path.Join(*(o.baseCfg.PathPrefix), PathPrefix, filePath)
	for i := 0; i < *o.baseCfg.RetryTime; i++ {
		_, err := o.service.PutObject(&s3.PutObjectInput{
			Bucket:        aws.String(*(o.baseCfg.Bucket)),
			Key:           aws.String(fullKey),
			ContentLength: aws.Int64((int64)(len(data))),
			ContentType:   aws.String(http.DetectContentType(data)),
			Body:          bytes.NewReader(data),
		})
		if err != nil {
			if i+1 == *o.baseCfg.RetryTime {
				return "", err
			}
			time.Sleep(time.Second * 2)
			continue
		}
		if o.baseCfg.Term != nil && *o.baseCfg.Term != "" {
			out, err := o.service.PutObjectTagging(&s3.PutObjectTaggingInput{
				Bucket: aws.String(*(o.baseCfg.Bucket)),
				Key:    aws.String(fullKey),
				Tagging: &s3.Tagging{
					TagSet: []*s3.Tag{
						{
							Key:   aws.String("term"),
							Value: o.baseCfg.Term,
						},
					},
				},
			})
			if err != nil {
				fmt.Println("PutObjectTagging", out, err)
			}
		}
		break
	}

	req, _ := o.service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(*(o.baseCfg.Bucket)),
		Key:    aws.String(fullKey),
	})

	// replace Scheme & Host
	if o.baseCfg.OutEndpoint != nil && *o.baseCfg.OutEndpoint != "" {
		value, err := url.Parse(*(o.baseCfg.OutEndpoint))
		if err != nil {
			return "", err
		}
		req.HTTPRequest.URL.Scheme = value.Scheme
		if value.Scheme == "" {
			req.HTTPRequest.URL.Scheme = "http"
		}
		req.HTTPRequest.URL.Host = value.Host
		if value.Path != "" {
			req.HTTPRequest.URL.Path = value.Path + req.HTTPRequest.URL.Path
		}
	}

	expire := 1 * time.Second
	if o.baseCfg.Expire != nil && *o.baseCfg.Expire > 0 {
		expire = time.Second * time.Duration(*o.baseCfg.Expire)
		urlStr, err = req.Presign(expire)
	} else {
		urlStr, err = req.Presign(expire)
		urlStr = strings.Split(urlStr, "?")[0]
	}

	return
}

func (o *FileServer) SaveKeepForever(key string, data []byte) (url string, err error) {
	return o.Save(key, data)
}

// Get 一个http url
func (o *FileServer) Get(ult string) ([]byte, error) {
	u, err := url.Parse(ult)
	if err != nil {
		return nil, err
	}
	p := u.Path
	p = strings.TrimLeft(p, "/")

	lis := strings.SplitN(p, "/", 2)
	out, err := o.service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(lis[0]),
		Key:    aws.String(lis[1]),
	})
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(out.Body)
	log.Printf("bucket and path is {{%s}} {{%s}},body size is %d kb", lis[0], lis[1], len(body)/1000)
	return body, err
}

type FileServerList struct {
	serverList []IS3ServerClient
}

func NewS3FileServer(config *FileServerConfig) (IS3ServerClient, error) {
	baseCfg := _DefaultS3FileServerConfig
	baseCfg.Merge(*config)
	endPoints := strings.Split(*baseCfg.Endpoint, ",")
	var outEndPoints []string
	if baseCfg.OutEndpoint != nil {
		outEndPoints = strings.Split(*baseCfg.OutEndpoint, ",")
	}
	for index := range endPoints {
		if len(outEndPoints) <= index {
			outEndPoints = append(outEndPoints, "")
		}
	}

	if len(endPoints) == 1 {
		baseCfg.OutEndpoint = &outEndPoints[0]
		return NewSingleS3Fileserver(&baseCfg)
	}

	retv := &FileServerList{}
	for i := range endPoints {
		baseCfg.Endpoint = &endPoints[i]
		baseCfg.OutEndpoint = &outEndPoints[i]
		clineTmp, err := NewSingleS3Fileserver(&baseCfg)
		if err != nil {
			return nil, err
		}
		retv.serverList = append(retv.serverList, clineTmp)
	}
	return retv, nil
}

func (o *FileServerList) hash(key string) int {
	hash := 0
	for _, v := range key {
		hash += int(v)
	}
	hash = hash % len(o.serverList)
	return hash
}

func (o *FileServerList) Init() error {
	for _, server := range o.serverList {
		if err := server.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (o *FileServerList) Save(key string, data []byte) (urlStr string, err error) {
	return o.serverList[o.hash(key)].Save(key, data)
}

func (o *FileServerList) SaveKeepForever(key string, data []byte) (url string, err error) {
	return o.serverList[o.hash(key)].SaveKeepForever(key, data)
}

// Get 一个http url
func (o *FileServerList) Get(ult string) ([]byte, error) {
	return o.serverList[o.hash(ult)].Get(ult)
}

func (c *FileServerConfig) Merge(other FileServerConfig) {
	var merged FileServerConfig
	ja, _ := json.Marshal(c)
	_ = json.Unmarshal(ja, &merged)
	jb, _ := json.Marshal(other)
	_ = json.Unmarshal(jb, &merged)
	*c = merged
}
