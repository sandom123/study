package study

import "fmt"

func main() {
	var a int
	var b int32
	a = 30
	b = int32(a) + int32(a)
	b = b + 5

	fmt.Println(a, b)
	fmt.Println("-----------------")

	s := "hello"
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(c, s2)
}
