package main

func main() {
	//printMsg("Calling a function")
	//
	//showMsg := func(msg string) {
	//	fmt.Println(msg)
	//}
	//showMsg("Second function")
	//
	//func(msg string) {
	//	fmt.Println(msg)
	//}("quick message")
	//
	//factor := 2
	//
	//mult := func(i int, j int) int {
	//	return i * j * factor
	//}
	//fmt.Println(mult(3, 4))

	////////////////////////
	//h1 := sayGreetings("ESP")
	//fmt.Println(h1())
	//h1 = sayGreetings("GER")
	//fmt.Println(sayGreetings("ENG")())
	//fmt.Println(h1())

	//num := getPositiveInt()
	//fmt.Println(num()) //1
	//fmt.Println(num()) //2
	//fmt.Println(num()) //3
	//
	//num2 := getPositiveInt()
	//fmt.Println(num2()) //1

	//addCounter, multCounter := addBy(), multBy()
	//
	//fmt.Println(addCounter(2))
	//fmt.Println(addCounter(3))
	//fmt.Println(multCounter(3))
	//fmt.Println(multCounter(3))

	/////////Callback function/////////

	//sqrt := func(i int) int {
	//	return i * i
	//}
	//cube := func(i int) int {
	//	return i * i * i
	//}
	//
	//fmt.Println(calc(sqrt, 3))
	//fmt.Println(calc(cube, 3))
	//
	//fmt.Println(calc(func(i int) int {
	//	return i * i
	//}, 3))

	////////////Assignment//////////

	//add := func(i, j int) int {
	//	return i + j
	//}
	//
	//mult := func(i, j int) int {
	//	return i * j
	//}
	//
	//fmt.Println(calc(add, 2, 4))
	//fmt.Println(calc(mult, 2, 4))
	//
	//fmt.Println(calc(func(i int, i2 int) int {
	//	return i + i2
	//}, 2, 4))
	//
	//fmt.Println(calc(func(i int, i2 int) int {
	//	return i * i2
	//}, 2, 4))

	//add := func(nums ...int) int {
	//	total := 0
	//	for _, i := range nums {
	//		total += i
	//	}
	//	return total
	//}
	//
	//mult := func(nums ...int) int {
	//	total := 1
	//	for _, i := range nums {
	//		total *= i
	//	}
	//	return total
	//}
	//fmt.Println(add(1, 2, 3, 4))
	//nums := []int{1, 2, 3, 4}
	//fmt.Println(add(nums...))
	//
	//fmt.Println(calc(add, nums))
	//fmt.Println(calc(mult, nums))

	////////////Factorial//////////
	//fmt.Println("=>", factorial(3))
	//fmt.Println("=>", factorial(4))
	//fmt.Println("=>", factorial(5))

	////////////Fibonacchi//////////
	//for i := 0; i <= 15; i++ {
	//	fmt.Printf("%d ", fibonacci(i))
	//}
	//fmt.Println(fibonacci(4))
	//defer showMsg()
	//fmt.Println("LOC 1", time.Now())
	//time.Sleep(5 * time.Second)
	//fmt.Println("LOC 2", time.Now())

	//fmt.Println(square(2))
	//fmt.Println(square(4))
	//fmt.Println(square(3))
}

//func square(x int) (result int) {
//	result = x * x
//
//	defer func() {
//		if x == 2 || x == 4 {
//			result += x
//		}
//	}()
//	fmt.Print("* ")
//	return
//}

//func showMsg() {
//	fmt.Println("\nSHOWMSG", time.Now())
//
//}

//func printMsg(msg string) {
//	fmt.Println(msg)
//}

//func sayGreetings(lang string) func() string {
//	h1 := func() string {
//		return "Hello"
//	}
//
//	if lang == "ESP" {
//		h1 = func() string {
//			return "Hola"
//		}
//	} else if lang == "GER" {
//		h1 = func() string {
//			return "Hallo"
//		}
//	}
//
//	return h1
//}

//func getPositiveInt() func() int {
//	i := 0
//	return func() int {
//		i++
//		return i
//	}
//}

//func addBy() func(int) int {
//	total := 0
//	return func(i int) int {
//		total += i
//		return total
//	}
//}
//
//func multBy() func(int) int {
//	total := 1
//	return func(i int) int {
//		total *= i
//		return total
//	}
//}

//func calc(f func(int) int, x int) int {
//	return f(x)
//}

////////////Assignment//////////

//func calc(f func(int, int) int, x, y int) int {
//	return f(x, y)
//}

//func calc(f func(...int) int, i []int) int {
//	return f(i...)
//}

//func factorial(n int) int {
//
//	if n == 0 {
//		return 1
//	}
//	fmt.Print(n, " ")
//
//	return n * factorial((n - 1))
//}

//func fibonacci(n int) int {
//	if n == 0 {
//		return 0
//	} else if n == 1 {
//		return 1
//	}
//	return fibonacci(n-1) + fibonacci(n-2)
//}
