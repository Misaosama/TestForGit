package main

import "fmt"

const (
    x = iota // x == 0
    y = iota // y == 1
    z = iota // z == 2
    w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说 w = iota，因此 w == 3。其实上面 y 和 z 可同样不用 "= iota"
)

const v = iota // 每遇到一个 const 关键字，iota 就会重置，此时 v == 0

const (
    h, i, j = iota, iota, iota // h=0,i=0,j=0 iota在同一行值相同
)

const (
    a       = iota // a=0
    b       = "B"
    c       = iota             //c=2
    d, e, f = iota, iota, iota //d=3,e=3,f=3
    g       = iota             //g = 4
)

func main(){
	// var num int = 9999
	// com := 9+9i
	// fmt.Println(num)
	// fmt.Println(com)

	// s := "hello"
	// fmt.Println(s)
	// c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	// fmt.Println(c)
	// c[0] = 'c'
	// fmt.Println(c)
	// s2 := string(c)  // 再转换回 string 类型
	// fmt.Printf("%s\n", s2)
	c := 9

	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

}