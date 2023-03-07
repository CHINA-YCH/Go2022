package main

import (
	"fmt"
	"math/rand"
)

/*
 * @Author: veno
 * @File: main
 * @Version: ...
 * @Date: 2023-03-06 19:49:05
 * @Description: ...
 */

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func main() {
	// 头部结构体
	var head Student
	head.Name = "head"
	head.Age = 28
	head.Score = 88

	// 尾部插入
	//TailInsert(&head)
	//Req(&head)

	// 头部插入
	Req(HeadInsert(&head))
}

// HeadInsert 头部插入
func HeadInsert(p *Student) *Student {
	for i := 0; i < 10; i++ {
		var stu = Student{
			Name:  fmt.Sprintf("stu%v", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		// 当前节点指向head，因为head是下一个节点
		stu.next = p // 指向下一个节点
		p = &stu     // 把当前的结构体给tail，让其继续循环
	}
	return p
}

// TailInsert 添加结构节点
func TailInsert(tail *Student) {
	for i := 0; i < 10; i++ {
		var stu Student = Student{
			Name:  fmt.Sprintf("stu%v", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
}

// Req 循环遍历
func Req(tmp *Student) {
	for tmp != nil {
		fmt.Println(*tmp)
		tmp = tmp.next
	}
}
