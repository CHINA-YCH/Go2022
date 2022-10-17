package main

import "fmt"

/*
一、接口的意义
好了，现在interface已经基本了解，那么接口的意义最终在哪里呢，想必 现在你已经有了一个初步的认识，实际上接口
的最大的意义就是实现多态的思想，就是我们可以根据interface类型来设计API接口，那么这种API接口的适应能力不仅能适应当下所实现的全部模块
也适应未来实现的模块来进行调用。
调用未来 可能就是接口的最大意义所在吧，这也是为什么架构师那么值钱，因为良好的架构师是可以针对interface设计一套框架，在未来许多年依然适用

二、依赖倒转原则
耦合度极高的模块关系设计
如下
我们来看上面的代码和图中每个模块之间的依赖关系，实际上并没有用到任何的interface接口层的代码，
显然最后我们的两个业务 张三开奔驰, 李四开宝马，程序中也都实现了。
但是这种设计的问题就在于，小规模没什么问题，但是一旦程序需要扩展，
比如我现在要增加一个丰田汽车 或者 司机王五， 那么模块和模块的依赖关系将成指数级递增，想蜘蛛网一样越来越难维护和捋顺。

三、优化
如上图，如果我们在设计一个系统的时候，将模块分为3个层次，抽象层、实现层、业务逻辑层。那么，
我们首先将抽象层的模块和接口定义出来，这里就需要了interface接口的设计，然后我们依照抽象层，依次实现每个实现层的模块，在我们写实现层代码的时候，
实际上我们只需要参考对应的抽象层实现就好了，实现每个模块，也和其它的实现的模块没有关系，这样也符合了上门介绍的开闭原则。
这样实现起来每个模块只依赖对象的接口，而和其它模块没关系，依赖关系单一。系统容易扩展和维护。

我们在指定业务逻辑也是一样，只需要参考抽象层的接口来业务就好了，抽象层暴雷出来的接口就是我们业务层可以使用的方法，然后可以通过多态的线下，接口指针指向哪个实现模块
调用来就是具体的实现方法，这样我们业务逻辑层也是依赖抽象成编程。

*/

type Audi struct {
}

func (a *Audi) Run() {
	fmt.Println("audi is running...")
}

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

type Zhang3 struct {
}

func (z *Zhang3) DriveBenz(benz *Benz) {
	fmt.Println("zhang3 Driver Benz")
	benz.Run()
}

func (z *Zhang3) DriverAudi(audi *Audi) {
	fmt.Println("zhang3 Driver Audi")
	audi.Run()
}

type Li4 struct {
}

func (l *Li4) DriverBenz(benz *Benz) {
	fmt.Println("li4 Driver Benz")
	benz.Run()
}

func (l *Li4) DriverAudi(audi *Audi) {
	fmt.Println("li4 Driver Audi")
	audi.Run()
}

func main() {
	// 业务1 张3开奔驰
	benz := &Benz{}
	zhang3 := Zhang3{}
	zhang3.DriveBenz(benz)

	// 业务2 李4开奥迪
	audi := &Audi{}
	li4 := &Li4{}
	li4.DriverAudi(audi)
}
