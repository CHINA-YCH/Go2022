package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main2
 * @Version: ...
 * @Date: 2022-10-27 16:21:13
 */
func main() {
	cooker := new(Cooker)
	ck := CommandCookChicken{cooker}
	ch := CommandCookChuaner{cooker}
	mm := new(WaiterMM)
	mm.CmdList = append(mm.CmdList, &ck)
	mm.CmdList = append(mm.CmdList, &ch)
	mm.Notify()

}

type Cooker struct {
}

func (c *Cooker) MakeChicken() {
	fmt.Println("烤串师傅烤了鸡肉串儿")
}

func (c *Cooker) MakeChuaner() {
	fmt.Println("烤串师傅烤了羊肉串儿")
}

// 抽象命令

type CommandE interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeChicken()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (ccc *CommandCookChuaner) Make() {
	ccc.cooker.MakeChuaner()
}

type WaiterMM struct {
	CmdList []CommandE
}

func (w *WaiterMM) Notify() {
	if w.CmdList == nil {
		return
	}
	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}
