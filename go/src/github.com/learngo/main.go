package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/learngo/something"
)

////////////////////////////////////////////////////////////

type person struct {
	name    string
	age     int
	favFood []string
}

////////////////////////////////////////////////////////////

var programName string = "testGo"

func main() {

	name := "jamesj"
	//const name string = "jamesj"

	fmt.Println("[" + programName + "] name: " + name)

	////////////////////////////////////////////////////////////

	something.SayHello()
	something.SayBye()

	////////////////////////////////////////////////////////////

	len, _ := lenAndUpper(name)
	fmt.Println("Len: " + strconv.Itoa(len) + ", upperName: ")

	////////////////////////////////////////////////////////////

	len2, uppercase := lenAndUpper2(name)
	fmt.Println("Len: " + strconv.Itoa(len2) + ", upperName: " + uppercase)

	////////////////////////////////////////////////////////////

	repeateMe("abc", "def", "ghi", "xyz")

	////////////////////////////////////////////////////////////

	len3, _ := lenAndUpper3(name)
	fmt.Println("Len: " + strconv.Itoa(len3))

	////////////////////////////////////////////////////////////

	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
	superAdd2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	////////////////////////////////////////////////////////////

	fmt.Println(canIDrink(16))

	////////////////////////////////////////////////////////////

	fmt.Println(canIDrink2(16))

	////////////////////////////////////////////////////////////

	pointer1()

	////////////////////////////////////////////////////////////

	array1()
	slice1()

	////////////////////////////////////////////////////////////

	map1()

	////////////////////////////////////////////////////////////

	struct1()

	////////////////////////////////////////////////////////////

	

	////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////
	// if len(os.Args) != 3 {
	// 	fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
	// 	os.Exit(1)
	// }

	// processType := os.Args[1]
	// ip := os.Args[2]

	// isWrongIp := UdpManager.CheckIp(ip)
	// if isWrongIp == false {
	// 	os.Exit(1)
	// }

	// if processType == "server" {
	// 	UdpManager.StartServer(ip)
	// } else if processType == "client" {
	// 	UdpManager.StartClient(ip)
	// }

}

// Multiple return
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 가변인자
func repeateMe(words ...string) {
	fmt.Println(words)
}

// 명시적 Multiple return
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// Defer > 함수 끝날 때 동작 정의
func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done.")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// for + Range
func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

// for + Range
func superAdd2(numbers ...int) {
	fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}

// Variable expression: If 문 안에 변수 선언해서 사용
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else if europeAge := age + 3; europeAge < 18 {
		return false
	}

	return true
}

// switch
func canIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 10:
		return true
	}

	// switch {
	// case age < 18:
	// 	return false
	// case age >= 18:
	// 	return true
	// }

	return false
}

// pointer
func pointer1() {
	a := 2
	b := &a

	a = 5

	fmt.Println(a, b)  // value, reference of a
	fmt.Println(a, *b) // value, value of a

	*b = 10

	fmt.Println(a, b)  // value, reference of a
	fmt.Println(a, *b) // value, value of a
}

// array
func array1() {
	names := [5]string{"a", "b", "c", "d", "e"}
	names[1] = "alala"
	fmt.Println(names)
}

// slice
func slice1() {
	names := []string{}
	//names[0] = "a"
	names = append(names, "a")
	names = append(names, "b")
	fmt.Println(names)
}

// map
func map1() {
	map1 := map[string]string{"a": "b"}

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	for key, _ := range map1 {
		fmt.Println(key)
	}
}

func struct1() {
	favFood := []string{"kimchi", "ramen"}
	person1 := person{"jamesj", 28, favFood}
	fmt.Println(person1)
}
