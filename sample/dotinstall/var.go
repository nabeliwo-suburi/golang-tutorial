// 変数
// 変数名： 1文字目に注意

package main

import "fmt"

func main2() {
	// var msg string
	// msg = "hello go"

	// var msg = "hello go"

	msg := "hello go"
	fmt.Println(msg)

	// var a, b int
	// a, b = 10, 15
	a, b := 10, 15

	var (
		c int
		d string
	)
	c = 20
	d = "hoge"
}
