// 基本的なデータ型
/*
string	"hello"
int		53
float	10.2
bool	false true
nil

初期値が決まっている
var s string // ""
var a int // 0
var f bool // false
*/
package main

import "fmt"

func main() {
	a := 10
	b := 12.3
	c := "hoge"
	var d bool
	fmt.Printf("a: %d, b: %f, c: %s, d: %t\n", a, b, c, d) // %dだと良い感じにする
}
