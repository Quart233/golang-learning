package pointer

import (
	"fmt"
	"reflect"
)

// Games - Games struct
type Games struct {
	name   string
	author string
}

// StructPointer - 声明指向结构体的指针
func StructPointer() {
	fmt.Println("++++++++ 声明指向结构体的指针 ++++++++++")
	whiteAblum := Games{
		name:   "白色相簿",
		author: "原田宇陀儿",
	}

	fmt.Println("whiteAblum", whiteAblum)

	var wP *Games
	fmt.Println("+++++ 输出变量类型 +++++++")
	fmt.Println(reflect.TypeOf(wP))

	wP = &whiteAblum

	fmt.Println("wP", wP)
	fmt.Println("wP", *wP)

}
