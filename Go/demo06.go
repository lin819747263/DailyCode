package main

import "fmt"

//全局变量
var a int =20

func main() {
	//局部变量
	var a int =10
	var b int =20
	var c int =0
	fmt.Printf("Main函数中a= %d\n",a)
	c=sum(a,b)
	fmt.Printf("Main函数中c= %d\n",c)
}
//形参变量
func sum(a,b int) int{
	fmt.Printf("sum函数中a的值为%d\n",a)
	fmt.Printf("sum函数中b的值为%d\n",b)

	return a+b;
}