package model

import (
	"fmt"
	"testing"
)

func TestGetFile(t *testing.T) {
	content, err := GetFile("telegraf.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func TestGetDir(t *testing.T) {
	filesInfo, err := GetDir(".")
	if err != nil {
		panic(err)
	}
	for _, v := range filesInfo {
		fmt.Println(v.Path)
	}
}
func TestUpdate(t *testing.T) {
	err := Update("telegraf.conf",
		"abc")
	if err != nil {
		panic(err)
	}
}
func TestTouch(t *testing.T) {
	err := Touch("telegraf.conf")
	if err != nil {
		panic(err)
	}
}
func TestDelete(t *testing.T) {
	err := Delete("telegraf.conf")
	if err != nil {
		panic(err)
	}
}
func TestExist(t *testing.T) {
	err := Exist("telegraf.conf")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("exist")
	}
}
