package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet2(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //封装函数内部的错误，传出给调用者
		return
	}
	//fmt.Println("resp = ", resp)
	//fmt.Printf("type resp: %T\n", resp)
	//fmt.Println("resp = ", resp.Body)
	//fmt.Printf("type resp: %T\n", resp.Body)

	defer resp.Body.Close()
	//循环读取网页数据，传递给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		fmt.Println(n, err)
		if n == 0 {
			fmt.Println("读取网页完成！")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}

		// 累加每次循环读取到的buf数据，存入result中
		result += string(buf[:n])
	}
	return
}

// 爬取单个页面的函数
func SpiderPage(i int, page chan<- int) {
	url := "https://tieba.baidu.com/f?kw=%E6%9D%8E%E6%AF%85&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := HttpGet2(url)
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}

	//fmt.Println("result = ", result)
	//将读到的整个网页存成一个文件
	f, err := os.Create("第 " + strconv.Itoa(i) + " 页.html")
	if err != nil {
		fmt.Println("Creat err", err)
		return
	}
	f.WriteString(result)
	f.Close() // 保存好一个文件，就关闭一个文件。

	page <- i //与主goroutine完成同步
}

// 爬取页面的操作
func working2(start, end int) {
	fmt.Printf("正在爬取第%d页到第%d页...\n", start, end)

	page := make(chan int)

	// 循环爬每一页的数据
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d网页完成爬取\n", <-page)

	}
}

func main() {
	// 指定爬取的起始、终止页
	var start, end int
	fmt.Print("请输入需要爬取的其实页面(>1): ")
	fmt.Scan(&start)

	fmt.Print("请输入终止页面(>start): ")
	fmt.Scan(&end)

	working2(start, end)
}
