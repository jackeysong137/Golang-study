package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//写入
	fileName := "./test.txt"
	//wirteFile(fileName)
	//追加写入
	//appendFile(fileName)
	writeBuf(fileName)
}

/*
清空并 写入
使用buf来实现写入
*/
func writeBuf(fileName string) {
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()
	writter := bufio.NewWriter(file)
	writter.WriteString("测试buf写入")
	writter.Flush() //刷入磁盘
}
func wirteFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close() //关闭文件
	file.WriteString("测试写入数据\n")
}
func appendFile(fileName string) {
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	file.WriteString("测试追加写入\n")
}
