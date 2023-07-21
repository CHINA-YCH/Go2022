package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
)

/*
 * @Author: veno
 * @File: main
 * @Version: ...
 * @Date: 2023-04-13 11:09:13
 * @Description: ...
 */

type Car struct {
	id          string
	tunnelID    string
	startFrame  int
	lastFrame   int
	grid        int
	gridHistory []int
	gridFrameID []int
	laneID      int
	laneHistory []int
	idSequence  []string
	position    int
	speed       float64
	features    [][]float32 // 根据实际情况选择合适的数据类型
	feaFrameID  []int
	bbox        []float64
}

func main() {
	//SetTesting()
	//ts2()
	//ts3()
	//ts4()
	//ts5()
	ts7()
}

func ts8() {

}

func ts7() {
	myMap := map[int64]float64{
		3: 10.0,
		1: 20.0,
		2: 5.0,
		8: 5.0,
	}
	v := sortAsc(myMap)
	log.Infof("%v", v)
}
func sortAsc(myMap map[int64]float64) []int64 {
	// 创建一个切片来保存map的key
	keys := make([]int64, 0, len(myMap))

	// 将map的key添加到切片中
	for k := range myMap {
		keys = append(keys, k)
	}

	// 使用sort.Slice函数按照map的value进行升序排序
	sort.Slice(keys, func(i, j int) bool {
		return myMap[keys[i]] <= myMap[keys[j]]
	})

	// 返回排序后的key
	var rs []int64
	for _, key := range keys {
		rs = append(rs, key)
	}
	return rs
}

func FindMaxIndex(numbers []float64) float64 {
	if len(numbers) == 0.0 {
		return -1
	}
	maxValue := numbers[0]
	maxIndex := 0
	for i, value := range numbers {
		if value > maxValue {
			maxIndex = i
			maxValue = value
		}
	}
	if float64(maxIndex) == 3 {
		println(true)
	}
	return float64(maxIndex)
}

func ts6() {
	numbers := []float64{3.14, 1.0, 2.5, 4.9, 2.2}
	index := FindMaxIndex(numbers)
	println(index)
}
func ts5() {
	// 创建四边形的顶点坐标
	quadrilateral := Quadrilateral{
		A: Point{0, 0},
		B: Point{1, 3},
		C: Point{3, 3},
		D: Point{2, 0},
	}

	// 创建待测试的点的坐标
	point := Point{1, 3.1}

	// 调用点是否在四边形内部函数，并获取到边的最短距离
	inside, distance := pointInsideQuadrilateral(point, quadrilateral)

	// 输出结果
	if inside {
		fmt.Println("点在四边形内部")
	} else {
		fmt.Println("点在四边形外部")
	}
	fmt.Printf("到边的最短距离: %.2f\n", distance)
}

func pointInsideQuadrilateral(p Point, q Quadrilateral) (bool, float64) {
	// 判断点是否在四边形内部，并计算点到边的最短距离
	inside := false
	minDistance := math.Inf(1)

	// 判断点是否在四边形的内部
	if pointInsideTriangle(p, q.A, q.B, q.C) || pointInsideTriangle(p, q.A, q.C, q.D) {
		inside = true
	}

	// 计算点到边的最短距离
	distances := []float64{
		pointToLineDistance(p, q.A, q.B),
		pointToLineDistance(p, q.B, q.C),
		pointToLineDistance(p, q.C, q.D),
		pointToLineDistance(p, q.D, q.A),
	}

	for _, dist := range distances {
		if dist < minDistance {
			minDistance = dist
		}
	}

	return inside, minDistance
}

func pointInsideTriangle(p, a, b, c Point) bool {
	// 判断点是否在三角形内部
	ab := (b.X-a.X)*(p.Y-a.Y) - (b.Y-a.Y)*(p.X-a.X)
	bc := (c.X-b.X)*(p.Y-b.Y) - (c.Y-b.Y)*(p.X-b.X)
	ca := (a.X-c.X)*(p.Y-c.Y) - (a.Y-c.Y)*(p.X-c.X)
	return (ab >= 0 && bc >= 0 && ca >= 0) || (ab <= 0 && bc <= 0 && ca <= 0)
}

type Point struct {
	X, Y float64
}

type Quadrilateral struct {
	A, B, C, D Point
}

func pointToLineDistance(p, v1, v2 Point) float64 {
	// 计算点到线段的最短距离
	a := v2.Y - v1.Y
	b := v1.X - v2.X
	c := v2.X*v1.Y - v1.X*v2.Y

	denom := math.Sqrt(a*a + b*b)
	if denom == 0 {
		return 0
	}

	//distance := math.Abs(a*p.X+b*p.Y+c) / denom
	distance := a*p.X + b*p.Y + c/denom
	return distance
}

//func ts4() {
//	// 创建四边形的顶点坐标
//	quadrilateral := []r3.Vector{
//		{0, 0, 0},
//		{1, 0, 0},
//		{1, 1, 0},
//		{0, 1, 0},
//	}
//
//	// 创建待测试的点的坐标
//	point := r3.Vector{0.5, 0.5, 0}
//
//	// 调用点与四边形的位置关系函数
//	position, distance := pointPolygonTest(quadrilateral, point)
//
//	// 输出结果
//	switch position {
//	case -1:
//		fmt.Println("点在四边形外部")
//	case 0:
//		fmt.Println("点在四边形边界上")
//	case 1:
//		fmt.Println("点在四边形内部")
//	}
//	fmt.Printf("到边的最近距离: %.2f\n", distance)
//
//}
//func pointPolygonTest(quadrilateral []r3.Vector, point r3.Vector) (int, float64) {
//	// 计算点与四边形的位置关系和到边的最近距离
//	intersects := false
//	onEdge := false
//	minDistance := math.Inf(1)
//
//	for i, j := 0, len(quadrilateral)-1; i < len(quadrilateral); i, j = i+1, i {
//		vi := quadrilateral[i]
//		vj := quadrilateral[j]
//
//		if (vi.Z > point.Z) != (vj.Z > point.Z) &&
//			point.X < (vj.X-vi.X)*(point.Z-vi.Z)/(vj.Z-vi.Z)+vi.X {
//			intersects = !intersects
//		}
//
//		if (vi.Z == point.Z && vj.Z == point.Z) && ((point.X >= vi.X && point.X <= vj.X) || (point.X >= vj.X && point.X <= vi.X)) {
//			onEdge = true
//			break
//		}
//
//		// 计算点到边的距离
//		distance := pointToLineDistance(point, vi, vj)
//		if distance < minDistance {
//			minDistance = distance
//		}
//	}
//
//	if onEdge {
//		return 0, minDistance
//	} else if intersects {
//		return 1, minDistance
//	} else {
//		return -1, minDistance
//	}
//}

//func pointToLineDistance(p, v1, v2 r3.Vector) float64 {
//	// 计算点到线段的最短距离
//	v := v2.Sub(v1)
//	w := p.Sub(v1)
//
//	c1 := w.Dot(v)
//	if c1 <= 0 {
//		return p.Sub(v1).Norm()
//	}
//
//	c2 := v.Dot(v)
//	if c2 <= c1 {
//		return p.Sub(v2).Norm()
//	}
//
//	b := c1 / c2
//	pb := v1.Add(v.Mul(b))
//	return p.Sub(pb).Norm()
//}

func ts3() {
	var arr2 = make([][]int64, 0)
	t1 := []int64{1, 1, 1}
	t2 := []int64{2, 2, 2}
	t3 := []int64{3, 3, 3}
	arr2 = append(arr2, t1)
	arr2 = append(arr2, t2)
	arr2 = append(arr2, t3)
	log.Infof("%+v", arr2)
}

func ts2() {
	var global_ids_list = make(map[int][]int)
	global_ids_list[0] = []int{0, 1, 2, 3}
	global_ids_list[1] = []int{0, 1, 2, 3}
	global_ids_list[2] = []int{0, 1, 2, 3} //
	global_ids_list[3] = []int{0, 1, 2, 3}

	carid_list := global_ids_list[2]
	for i := 0; i < 4; i++ {
		x := carid_list[i]
		log.Infof("- - - - -%+v", x)
		if x != 2 {
			carid_list = global_ids_list[x]
			//global_ids_list[x] = carid_list
		}
	}

}
func SetTesting() {
	var mp = make(map[string]Car)
	mp["xx"] = Car{
		id: "110",
	}
	//log.Infof("mp:%+v", mp)

	car := mp["xx2"]
	car.id = "112"
	mp["xx2"] = car

	car3 := mp["xx3"]
	car3.id = "113"
	mp["xx3"] = car3

	//ls := mp["xx3"]

	for i := 0; i < 3; i++ {

	}
	//log.Infof("car:%+v", car)

	//log.Infof("mp:%+v", mp["xx"])

	//car2 := mp["xx"]
	//car2.id = "112"
	//car2.set()
	//mp["xx"] = car2
	//
	//log.Infof("mp---:%+v", mp["xx"])
	//
	//delete(mp, "110")
	//
	//log.Infof("mp:%+v", mp["xx"])

}

func (c *Car) set() {
	c.id = "113"
}

type User struct {
	Id   int
	Name string
	Age  int
}

func MarTesting() {
	mp := make(map[string]User)
	mp["u1"] = User{
		Id:   1,
		Name: "王一",
		Age:  18,
	}
	mp["u2"] = User{
		Id:   2,
		Name: "王二",
		Age:  19,
	}
	if v, ok := mp["u1"]; ok {
		v.Age = 188
		mp["u1"] = v
	}
	marshal, _ := json.Marshal(mp)
	println(string(marshal))
}
