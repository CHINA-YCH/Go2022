package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var i = 0

func main() {
	http.HandleFunc("/v1/facilities-manage-service/access_control/digraph", handler2)
	go func() {
		err := http.ListenAndServe("localhost:32000", nil)
		if err != nil {
			log.Errorf("listen and server error:%v", err)
		}
	}()
	log.Infof("- - - - - - -1111111 - - - -")
	time.Sleep(10000 * time.Second)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	s := "{\n    \"code\": 0,\n    \"message\": \"OK\",\n    \"data\": {\n        \"straight\": {\n            \"nodes\": [\n                {\n                    \"id\": 1,\n                    \"thirdCode\": \"A\",\n                    \"bayonetType\": 0,\n                    \"acType\": 1,\n                    \"pileNo\": 0\n                },\n                {\n                    \"id\": 2,\n                    \"thirdCode\": \"B\",\n                    \"bayonetType\": 1,\n                    \"acType\": 1,\n                    \"pileNo\": 1000\n                },\n                {\n                    \"id\": 3,\n                    \"thirdCode\": \"B2\",\n                    \"bayonetType\": 1,\n                    \"acType\": 2,\n                    \"pileNo\": 0\n                },\n                {\n                    \"id\": 4,\n                    \"thirdCode\": \"C\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 1000\n                },\n                {\n                    \"id\": 5,\n                    \"thirdCode\": \"D\",\n                    \"bayonetType\": 0,\n                    \"acType\": 2,\n                    \"pileNo\": 2000\n                },\n                {\n                    \"id\": 6,\n                    \"thirdCode\": \"E\",\n                    \"bayonetType\": 0,\n                    \"acType\": 1,\n                    \"pileNo\": 3000\n                },\n                {\n                    \"id\": 7,\n                    \"thirdCode\": \"F\",\n                    \"bayonetType\": 0,\n                    \"acType\": 2,\n                    \"pileNo\": 4000\n                },\n                {\n                    \"id\": 8,\n                    \"thirdCode\": \"E1\",\n                    \"bayonetType\": 1,\n                    \"acType\": 1,\n                    \"pileNo\": 4000\n                },\n                {\n                    \"id\": 9,\n                    \"thirdCode\": \"E2\",\n                    \"bayonetType\": 1,\n                    \"acType\": 2,\n                    \"pileNo\": 3000\n                }\n            ],\n            \"edges\": [\n                {\n                    \"from\": 1,\n                    \"to\": 4,\n                    \"distance\": 1000\n                },\n                {\n                    \"from\": 1,\n                    \"to\": 3,\n                    \"distance\": 100\n                },\n                {\n                    \"from\": 2,\n                    \"to\": 4,\n                    \"distance\": 150\n                },\n                {\n                    \"from\": 4,\n                    \"to\": 5,\n                    \"distance\": 1000\n                },\n                {\n                    \"from\": 6,\n                    \"to\": 7,\n                    \"distance\": 1000\n                },\n                {\n                    \"from\": 8,\n                    \"to\": 7,\n                    \"distance\": 200\n                },\n                {\n                    \"from\": 6,\n                    \"to\": 9,\n                    \"distance\": 150\n                }\n            ]\n        },\n        \"reverse\": {\n            \"nodes\": [\n                {\n                    \"id\": 107,\n                    \"thirdCode\": \"F1\",\n                    \"bayonetType\": 0,\n                    \"acType\": 2,\n                    \"pileNo\": 0\n                },\n                {\n                    \"id\": 106,\n                    \"thirdCode\": \"E1\",\n                    \"bayonetType\": 0,\n                    \"acType\": 1,\n                    \"pileNo\": 2000\n                },\n                {\n                    \"id\": 101,\n                    \"thirdCode\": \"A1\",\n                    \"bayonetType\": 0,\n                    \"acType\": 2,\n                    \"pileNo\": 3000\n                },\n                {\n                    \"id\": 104,\n                    \"thirdCode\": \"C1\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 3500\n                },\n                {\n                    \"id\": 105,\n                    \"thirdCode\": \"D1\",\n                    \"bayonetType\": 0,\n                    \"acType\": 1,\n                    \"pileNo\": 4000\n                },\n                {\n                    \"id\": 108,\n                    \"thirdCode\": \"C11\",\n                    \"bayonetType\": 1,\n                    \"acType\": 1,\n                    \"pileNo\": 3500\n                },\n                {\n                    \"id\": 109,\n                    \"thirdCode\": \"C12\",\n                    \"bayonetType\": 1,\n                    \"acType\": 2,\n                    \"pileNo\": 3000\n                }\n            ],\n            \"edges\": [\n                {\n                    \"from\": 104,\n                    \"to\": 101,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 105,\n                    \"to\": 104,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 106,\n                    \"to\": 107,\n                    \"distance\": 2000\n                },\n                {\n                    \"from\": 104,\n                    \"to\": 109,\n                    \"distance\": 200\n                },\n                {\n                    \"from\": 108,\n                    \"to\": 101,\n                    \"distance\": 100\n                }\n            ]\n        }\n    }\n}"
	i++
	fmt.Fprintf(w, "%v", s)
	log.Infof("exec function %v. . . . .", i)
	go test()
}

func test() {
	time.Sleep(3 * time.Second)
	log.Infof("- - - - - this is after sleep 3 second")
}
