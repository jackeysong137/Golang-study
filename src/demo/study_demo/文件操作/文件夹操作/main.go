package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	////创建文件夹 有就创建 没有就不创建
	err := os.MkdirAll("./dir1/dir2/dir3", 0777)
	if err != nil {
		fmt.Println(err)
	}
	// 遍历文件夹
	filepath.Walk("./dir1", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Printf("目录: %s\n", path)
		} else {
			fmt.Printf("文件: %s\n", path)

		}
		return nil
	})
	//删除 递归删除 os.RemoveAll() 删除单个os.Remove()

	err1 := os.RemoveAll("./dir1")
	if err1 != nil {
		fmt.Println(err1)
	}
}
