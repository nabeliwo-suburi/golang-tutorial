// 構造体
// メソッド(データ型に紐付いた関数) オブジェクト指向的なやつではない

package main

import "fmt"

type user struct {
	name  string
	score int
}

func (u user) show() { // レシーバー 通常は値渡しだが、u *userにすると参照渡しになる
	fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

func (u *user) hit() {
	u.score++
}

func main() {
	// u := new(user)
	// // (*u).name = "taguchi"
	// u.name = "taguchi"
	// u.score = 20

	// u := user{"taguchi", 50}
	u := user{name: "taguchi", score: 50}
	// fmt.Println(u)
	u.hit()
	u.show()
}
