package main

import "fmt"

// ocp（open-closed principle） 开闭原则，对扩展时开放的，对修改是关闭的

type Pet interface {
	eat()
	sleep()
}
type Dog struct {
}

type Cat struct {
}

func (dog *Dog) eat() {
	fmt.Printf("dog eat...\n")
}
func (dog *Dog) sleep() {
	fmt.Printf("dog sleep...\n")
}
func (cat *Cat) eat() {
	fmt.Printf("cat eat...\n")
}
func (cat *Cat) sleep() {
	fmt.Printf("cat sleep...\n")
}

type Person struct{}

func (person *Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}
func main() {
	var dog = Dog{}
	var dog2 = &dog
	var cat = Cat{}
	var cat2 = &cat
	var person = Person{}
	person.care(dog2)
	fmt.Printf("- - - - - - - - - - - - - - - - -\n")
	person.care(cat2)
}
