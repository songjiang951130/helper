package service

import "helper/api/danjuan"

type User struct {
	Name  string
	Email string
	Git   string
	Dan   []danjuan.Items
}

func GetIndexInfo() User {
	// u := new(User)
	// u.name = "宋江"
	// u.email = "songjiangcq@qq.com"
	// u.git = "https://gitee.com/songjiang_951130/helper.git"
	dan := GetLowIndex()
	// u := User{"宋江", "songjiangcq@qq.com", "https://gitee.com/songjiang_951130/helper.git"}
	u := User{"宋江", "songjiangcq@qq.com", "https://gitee.com/songjiang_951130/helper.git", dan}
	return u
}
