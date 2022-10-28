package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2022-10-27 11:29:00
 */
func main1() {
	doctor := new(Doctor)
	cmdEye := CommandTreadEye{doctor}
	cmdNose := CommandTreatNose{doctor}

	// 护士
	n := new(Nurse)
	n.CmdList = append(n.CmdList, &cmdEye)
	n.CmdList = append(n.CmdList, &cmdNose)

	// 执行
	n.Notify()
}

// 医生-命令接收者

type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 抽象的命令

type Command interface {
	Treat()
}

// 治疗眼睛的病单

type CommandTreadEye struct {
	doctor *Doctor
}

func (ct *CommandTreadEye) Treat() {
	ct.doctor.treatEye()
}

type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// 护士 - 调用命令者

type Nurse struct {
	CmdList []Command
}

// 发送病单，发送命令的方法

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat() // 制定病单绑定的命令
	}
}
