package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2023-02-01 18:22:04
 */

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"code\":0,\"message\":\"OK\",\"data\":{\"straight\":{\"nodes\":[{\"id\":2,\"thirdCode\":\"ty735985\",\"bayonetType\":0,\"acType\":1,\"pileNo\":735985},{\"id\":1,\"thirdCode\":\"ty737050\",\"bayonetType\":0,\"acType\":2,\"pileNo\":737050}],\"edges\":[{\"from\":2,\"to\":1,\"distance\":1065}]},\"reverse\":{\"nodes\":[{\"id\":4,\"thirdCode\":\"sjz737000\",\"bayonetType\":0,\"acType\":2,\"pileNo\":737000},{\"id\":3,\"thirdCode\":\"sjz735940\",\"bayonetType\":0,\"acType\":1,\"pileNo\":735940}],\"edges\":[{\"from\":4,\"to\":3,\"distance\":1060}]}}}")
}
func main() {
	http.HandleFunc("/v1/facilities-manage-service/access_control/digraph", sayHi)
	log.Fatal(http.ListenAndServe("localhost:18083", nil))
}
