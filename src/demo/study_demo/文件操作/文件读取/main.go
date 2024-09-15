package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
文件的读取 有三种方式
	1: os.open() 只读
	2:os.OpenFile 可以设置权限
	3:ioutil
*/

func main() {
	//os.open 读取
	fileName := "main.go"
	// fileContent := readFileByOpen(fileName)
	// fmt.Println(fileContent)

	// fileContent1 := readFileByOpenFile(fileName)
	// fmt.Println(fileContent1)
	// bufio 读取文件
	str := readFileByBuf(fileName)
	fmt.Println(str)

}
func readFileByBuf(fileName string) string {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close() //操作文件完成必须关闭文件流 ，否则会造成内存泄漏
	reader := bufio.NewReader(file)
	sb := bytes.NewBufferString("")
	for {
		str, err := reader.ReadString('\n') //每次读取一行
		if err != nil && err != io.EOF {
			return ""
		}
		//完结 使用buf读取的时候 读到末尾 可能是有字符串
		sb.WriteString(str)

		//完结 使用buf读取的时候 读到末尾 可能是有字符串
		if err == io.EOF {
			break
		}

	}
	return sb.String()
}

func readFileByOpenFile(fileName string) string {
	/*
		os.O_CREATE 不存在就会创建
		os.O_RDONLY 只读
		os.O_APPEND 追加
		os.O_WRONLY 读写
		os.O_TRUNC 清空
		多个模式 可以使用 os.O_CREATE ｜os.O_WRONLY ｜os.O_APPEND 这样使用
	*/

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666) //0666 有读写执行权限 755等 可以设置
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	//
	fileSlice := make([]byte, 100)
	sliceByte := make([]byte, 100)
	for {
		n, err := file.Read(sliceByte)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return ""
		}
		//读取完成
		if err == io.EOF {
			break
		}
		fileSlice = append(fileSlice, sliceByte[:n]...)
	}
	return string(fileSlice)

}

func readFileByOpen(fileName string) string {
	file, err := os.Open(fileName) //只读的方式打开
	//操作文件完成必须关闭文件流 ，否则会造成内存泄漏
	if err != nil {
		fmt.Println("读取文件失败:", err)
	}
	defer file.Close()
	fileSlice := make([]byte, 100)
	sliceByte := make([]byte, 100)

	for {
		n, err := file.Read(sliceByte)
		if err != nil && err != io.EOF {
			fmt.Println("读取文件失败:", err)
		}
		//读取完结  err == io.EOF
		if err == io.EOF {
			break
		}
		//追加有效的部分
		fileSlice = append(fileSlice, sliceByte[:n]...)
	}

	return string(fileSlice)
}
