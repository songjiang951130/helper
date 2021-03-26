package base

import "fmt"

type User struct {
	ID   int
	Name string
}

func PlayStruct() {
	p := new(User)
	obj := *p
	obj.ID = 123
	obj.Name = "张三"

	fmt.Println(p.ID, p.Name)
	fmt.Println(obj.ID, obj.Name)
}

// go test -v  ./base
// slice翻倍扩容
func PlaySlice() {
	arr := []int{1, 2, 3}
	//从下标数[1,2)
	s := arr[1:2]
	fmt.Println("copy s:", s)
	s[0] = 100
	for i := 20; i < 30; i++ {
		fmt.Printf("before len:%4d cap:%4d\n", len(s), cap(s))
		s = append(s, i)
		fmt.Printf("after  len:%4d cap:%4d\n", len(s), cap(s))
	}
	fmt.Println("arr:", arr)
	fmt.Println("s:", s)
}
