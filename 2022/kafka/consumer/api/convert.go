package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

/*
 * @Author: veno
 * @File: convert
 * @Version: ...
 * @Date: 2023-02-28 16:59:12
 * @Description: ...
 */

func convert(bytes []byte) *MsgInfoSource {
	var msg = MsgInfoSource{}
	err := json.Unmarshal(bytes, &msg)
	if err != nil {
		log.Errorf("format msg error:%+v", err)
	}
	return &msg
}
