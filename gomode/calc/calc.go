package calc //packageg 关键字定义包 ♥️ 必须在顶部  包名最好和文件夹的名称相同 可以不相同
/*
一个文件夹里面只能定义一个包名  ， 多个go文件的额包名是一样的 可以通过import导入
*/
import "fmt"

func Add(a, b int) int {
	return a + b
}

func init() {

	fmt.Println("calc init ...")
}
