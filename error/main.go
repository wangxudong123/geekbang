package main

import (
	"error/dao"
	"fmt"
)

func main() {
	db := dao.NewModel()
	if _, err := db.Select(); err != nil {
		fmt.Printf("%+v", err)
	}
}
