package main

import "fmt"

//goto的使用
func gotof(){
	i := 0
Here:
	println(i)
	i ++
	if i > 100{
		return
	}
	goto Here
}

func unhex(c byte) byte{
	switch  {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}

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

	gotof()
	fmt.Println(unhex('g'))
}
