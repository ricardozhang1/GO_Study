package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
)

// HttpGetPage 爬取指定的url页面，并返回result
func HttpGetPage(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	//循环获取整页数据
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成！")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}

	return
}


func spiderPage2(index int)  {
	//获取url
	url := "https://tieba.baidu.com/f?kw=%E6%9D%8E%E6%AF%85&ie=utf-8&pn=" + strconv.Itoa((index-1)*50)

	//封装HttpGet2获取页面
	ret, err := HttpGetPage(url)
	if err != nil {
		fmt.Println("HttpGetPage err", err)
		return
	}
	fmt.Println("result = ", ret)
}


func toWork(start, end int)  {
	fmt.Printf("正在爬取 %d 页到 %d 页...\n", start, end)

	for i := start; i <= end; i++ {
		spiderPage2(i)
	}
}


func main() {
	//指定爬取的起始页和终止页
	var start, end int
	fmt.Print("请输入起始页面(>=1): ")
	fmt.Scan(&start)
	fmt.Print("请输入结束页面(>=start): ")
	fmt.Scan(&end)

	toWork(start, end)
}
