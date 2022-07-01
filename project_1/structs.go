package main

import "fmt"

//type player struct {
//	name, sport string
//	age         int
//}

//type employee struct {
//	name, job    string
//	lastLoggedIn string
//	DOB          date.Date
//}

//type book struct {
//	name   string
//	author string
//}

//type player struct {
//	name, sport string
//	age         int
//	info        generalInfo
//}
//type generalInfo struct {
//	country, color string
//}

type movie struct {
	name, actor string
}

func (m movie) fullInfo() string {
	return m.name + " " + m.actor
}

func main() {

	m1 := movie{"Mike", "Tom kruz"}
	fmt.Println(m1.fullInfo())
	//pl := player{"21312", "12312312", 44, generalInfo{"2333", "55555"}}
	//fmt.Println(pl)

	//b1 := book{"Go Programming", "Author 1"}
	//b2 := book{"Go Programming", "Author 1"}
	//b3 := book{"C++ Programming", "Author 2"}
	//
	//fmt.Printf("LOC1 - %v\n", b1)
	//
	//fmt.Println(compare1(b1, b2))
	//fmt.Println(compare1(b1, b3))
	//
	//fmt.Println(compare2(&b1, &b2))
	//fmt.Printf("LOC3 - %v\n\n", b1)
	//
	//p1 := &book{"Go Web Programming", "Author 3"}
	//
	//p2 := new(book)
	//*p2 = book{"Go Web Programming", "Author 3"}
	//
	//fmt.Printf("LOC4 - &p1=%v *p1=%v\n", &p1, *p1)
	//
	//fmt.Println(compare2(p1, p2))
	//fmt.Printf("LOC5 - &p1=%v *p1=%v\n", &p1, *p1)

	//var emp employee
	//
	//emp.name = "Mike"
	//emp.job = "GO prog"
	//emp.lastLoggedIn = time.Now().Format(time.RFC850)
	//emp.DOB = date.Today()
	//fmt.Println(emp)
	//
	//p := &emp
	//p.name = "UUUUUU"
	//fmt.Println(*p)
	//fmt.Println(emp)

	//type myType float64
	//var total myType
	//
	//total = 44
	//fmt.Printf("%.2f %T \n", total, total)
	//
	//var total2 float64
	//
	//total2 = float64(total)
	//fmt.Printf("%.2f %T \n", total2, total2)

	//player1 := player{"Messi", "Scoorer", 24}
	//fmt.Println(player1)
	//player2 := player{
	//	name:  "2",
	//	sport: "2",
	//	age:   2,
	//}
	//fmt.Println(player2)
	//
	//var player3 player
	//fmt.Println(player3)
	//player3.name = "Mike"
	//player3.sport = "boxing"
	//player3.age = 2
	//fmt.Println(player3)
	//
	//player4 := player{
	//	name:  "SMTH",
	//	sport: "smht",
	//}
	//fmt.Println(player4)
	//
	//player5 := new(player)
	//player5.name = "UUU"
	//fmt.Println(*player5)

	//player1 := struct {
	//	name, sport string
	//	age         int
	//}{"Mike", "1212312", 3}
	//
	//fmt.Println(player1)

	//testSizeOfStruct()

}

//func compare1(b book, q book) bool {
//	if b.name == q.name && b.author == q.author {
//		b.name = "Java programming"
//		return true
//	}
//	return false
//}
//
//func compare2(b *book, q *book) bool {
//	if b.name == q.name && b.author == q.author {
//		b.name = "Python Programming"
//		return true
//	}
//	return false
//}

//func testSizeOfStruct() {
//	type values struct {
//		val1 int32 // 4 bytes
//		val2 int32 // 4 bytes
//		val3 int64 // 8 bytes
//		val4 int64 // 8 bytes
//	}
//
//	x := values{10, 20, 30, 40}
//	p := &x
//	fmt.Println(unsafe.Sizeof(p), unsafe.Sizeof(*p), unsafe.Sizeof(x))
//
//	q := &values{11, 21, 31, 41}
//	fmt.Println(unsafe.Sizeof(q), unsafe.Sizeof(*q))
//}
