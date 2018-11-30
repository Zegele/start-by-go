//当前程序的包名
package main //必须放在第一行

//导入其他的包
import "fmt" //导入的包必须被使用

//常量的定义
const PI = 3.141592653 //常量赋值

//全局变量的声明与赋值
var name = "gopher" //全局变量能在整个package中使用

//一般类型声明
type newType int //newType是一般类型的名称

//结构的声明
type gopher struct{} //必须由大括号，需要填充内容的。

//接口的声明
type golang interface{} //使用interface作为第二个关键字，与上面的struct一样

//由main函数作为程序入口点启动
func main() {
	fmt.Println("NiuBiLity!!!")
}
