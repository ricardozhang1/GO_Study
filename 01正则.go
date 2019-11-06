package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc cat 8ca azc cba"
	//解析、编译正则表达式
	ret := regexp.MustCompile(`a.c`)

	//提取需要的信息
	alls := ret.FindAllStringSubmatch(str, -1)
	fmt.Println("alls = ", alls)
}




