package xueqiu

import (
	"fmt"
	"os"
)

// const cookie = ""

func GetCookie() string {
	file, err := os.Open("cookie.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	file.Read(buffer)
	res := string(buffer)
	fmt.Println("bytestream to string: ", string(buffer))
	return res
}
