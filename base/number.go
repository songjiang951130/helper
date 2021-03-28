package base

import "fmt"

func Cacl() {
	rank := 5
	res := rank * 100.0 / 200
	fmt.Printf("%T", res)
	fmt.Println("res", res)
}
