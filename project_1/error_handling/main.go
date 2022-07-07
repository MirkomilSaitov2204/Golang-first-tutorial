package main

import (
	"errors"
	"fmt"
	"log"
)

//var fileName = "some.txt"

func main() {
	//testCase1()
	//testCase2()
	//testCase3()
	//fmt.Println("Control")

	_, _, err := produceError1(true)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Fatalln("NO error")
	}
}

func produceError1(b bool) (int, int, error) {
	if !b {
		return 0, 0, nil
	}
	errMsg1 := errors.New("I'm just a value like any other values in Go! (Rob Pike told me)")
	fmt.Printf("**errors.New() generates type: %T\n", errMsg1)

	return 0, 0, errMsg1
}

//func testCase1() {
//	_, err := os.Open(fileName)
//	if err != nil {
//		log.Println("Error", err)
//	}
//}
//
//func testCase2() {
//	_, err := os.Open(fileName)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
//
//func testCase3() {
//	_, err := os.Open(fileName)
//	if err != nil {
//		panic(err)
//	}
//}
