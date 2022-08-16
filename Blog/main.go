package main

import (
	"errors"
	"fmt"
)

//const numPool = 1000
//
//func CalcuateValue(intChan chan int) {
//	randomNumber := helpers.RandomNumber(numPool)
//	intChan <- randomNumber
//}

//type myStruct struct {
//	FirstName string
//}
//
//func (m *myStruct) printFirstName() string {
//	return m.FirstName
//}

//type User struct {
//	Firstname string
//	Lastname  string
//}

//type Animal interface {
//	Says() string
//	NumberOfLegs() int
//}
//
//type Dog struct {
//	Name  string
//	Breed string
//}
//
//type Gorilla struct {
//	Name          string
//	Color         string
//	NumberOfTeeth int
//}

func main() {
	result, err := devide(100.0, 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	//intChan := make(chan int)
	//defer close(intChan)
	//
	//go CalcuateValue(intChan)
	//
	//num := <-intChan
	//fmt.Println(num)
	//dog := Dog{
	//	Name:  "Simon",
	//	Breed: "German",
	//}
	//PrintInfo(dog)
	//
	//gorilla := Gorilla{
	//	Name:          "Ling kong",
	//	Color:         "Black",
	//	NumberOfTeeth: 32,
	//}
	//PrintInfo(gorilla)
	//myMap := make(map[string]User)
	//me := User{
	//	Firstname: "adasd",
	//	Lastname:  "asdgfggg",
	//}
	//
	//myMap["me"] = me
	//println(myMap["me"].Firstname)
	//var myVar myStruct
	//myVar.FirstName = "Hello"
	//myVar2 := myStruct{
	//	FirstName: "Hello world",
	//}
	//println(myVar.printFirstName())
	//println(myVar2.printFirstName())

	//var whatToSay string
	//
	//whatToSay, _ = saySomething("Hello world")
	//
	//fmt.Println(whatToSay)
	//fmt.Println(saySomething("Hello world"))

	//var myString string
	//myString = "Green"
	//
	//fmt.Println("my string is set to ", myString)
	//changeUsingPointer(&myString)
	//fmt.Println("my string is set to ", myString)

}

//func (d Dog) Says() string {
//	return "woof"
//}
//func (d Dog) NumberOfLegs() int {
//	return 4
//}
//
//func (d Gorilla) Says() string {
//	return "woof"
//}
//func (d Gorilla) NumberOfLegs() int {
//	return 4
//}
//
//func PrintInfo(a Animal) {
//	log.Fatal("This animal says ", a.Says(), " and it has ", a.NumberOfLegs(), " Legs")
//}

//func changeUsingPointer(s *string) {
//	newValue := "Red"
//	*s = newValue
//}

//func saySomething(s string) (string, string) {
//	return s, "World"
//}

func devide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("Can not devide by 0")
	}
	result = x / y
	return result, nil
}
