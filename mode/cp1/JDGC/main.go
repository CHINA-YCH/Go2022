package main

/*
为什么需要工厂模式
首先来看，如果没有工厂模式，在开发者创建一个类的对象时，如果有很多不同种类的对象将会如何实现，代码如下：
*/
func main() {

}

// 水果类

type Fruit struct {
}

func (f *Fruit) Show(name string) {

}
