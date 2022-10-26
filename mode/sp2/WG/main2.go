package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main2
 * @Version: ...
 * @Date: 2022-10-26 15:43:20
 */
func main() {
	homePlayer := new(HomePlayerFacade)
	homePlayer.DoKTV()
	fmt.Println("- - - - - - -")
	homePlayer.DoGame()
}

type TV struct {
}

func (t *TV) On() {
	fmt.Println("打开 电视机")
}
func (t *TV) Off() {
	fmt.Println("关闭 电视机")
}

type VoiceBox struct {
}

func (t *VoiceBox) On() {
	fmt.Println("打开 音箱")
}
func (t *VoiceBox) Off() {
	fmt.Println("关闭 音箱")
}

type Light struct {
}

func (t *Light) On() {
	fmt.Println("打开 灯光")
}
func (t *Light) Off() {
	fmt.Println("关闭 灯光")
}

type Xbox struct {
}

func (t *Xbox) On() {
	fmt.Println("打开 游戏机")
}
func (t *Xbox) Off() {
	fmt.Println("关闭 游戏机")
}

type MicroPhone struct {
}

func (m *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}

func (m *MicroPhone) Off() {
	fmt.Println("关闭 麦克风")
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("打开 投影仪")
}

func (p *Projector) Off() {
	fmt.Println("关闭 投影仪")
}

// 外观

type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	light Light
	xbox  Xbox
	mp    MicroPhone
	pro   Projector
}

// KTV模式

func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("家庭影院进入KTV模式")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.light.Off()
	hp.vb.On()
}

// 游戏模式

func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("家庭影院进入Game模式")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}
