package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//这只是一个简单的版本，只是获取QQ邮箱并没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

//HandleError 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

//GetEmail 爬邮箱
func GetEmail() {
	//1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get Error!")
	defer func() { _ = resp.Close }()

	//2. 读取页面
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll Error!")
	//字节转换为字符串
	pageContent := string(pageBytes)
	//fmt.Println(pageContent)
	//3.过滤数据，过滤QQ邮箱
	re := regexp.MustCompile(reQQEmail)
	//-1代表取全部
	results := re.FindAllStringSubmatch(pageContent, -1)
	//fmt.Println(results)

	//遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

func main() {
	GetEmail()
}
