package producer

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type MsgSink interface {
	Push(ctx context.Context, msg string) error
	PushList(ctx context.Context, msgs []string) error
}

type Exec interface {
	MsgProcess(msg string)
	MsgProcessTest(msg [][]byte, errNum *int64)
}

type exec struct {
	sinkInit MsgSink
}

func NewExec(sinkInit MsgSink) Exec {
	return &exec{
		sinkInit: sinkInit,
	}
}

func (e *exec) MsgProcessTest(msg [][]byte, _ *int64) {
	var batch []string
	for _, msg := range msg {
		var str string
		str = string(msg)
		batch = append(batch, str)
	}
	background := context.Background()
	i := len(batch)
	if i != 0 && batch != nil {
		var s = batch[i-1]
		batch[i-1] = fmt.Sprintf("%s - - - - - - - - - - - - ", s)
	}
	_ = e.sinkInit.PushList(background, batch)
}

func (e *exec) MsgProcess(msg string) {
	background := context.Background()
	_ = e.sinkInit.Push(background, msg)
	marshal, _ := json.Marshal(msg)
	log.Println("\n* * * * * * pushOutList: %s\n", string(marshal))
}
