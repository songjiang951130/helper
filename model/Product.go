package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Display() {
	fmt.Print("import llll")
}
