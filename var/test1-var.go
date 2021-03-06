package main

import "fmt"

// 声明全局变量，方法1，方法2，方法3都是可以的
var gA int 
var gB = 100
// 第四种声明方式只能够用在函数体内声明

// 四种变量声明方式
func main() {
	// 第一种：var默认声明的int如果不赋值则默认为0
	var a int
	fmt.Println("a=",a)
	fmt.Printf("\ntype of a = %T",a)

	// 第二种：声明一个变量，给一个初始化值
	var b int = 100
	fmt.Println("\nb=",b)
	fmt.Printf("\ntype of b = %T",b)

	var bb string = "asd"
	fmt.Printf("\ntype of bb = %T",bb)

	// 第三种：不声明类型，通过值直接给他赋值
	var c = 100
	fmt.Println("\nc=",c)
	fmt.Printf("\ntype of c = %T",c)
	
	var cc = "asd"
	fmt.Printf("\ntype of cc = %T",cc)

	// 第四种：省去var关键词，直接自动匹配
	d := 100
	fmt.Printf("\ntype of d = %T",d)

	e := 100
	fmt.Printf("\ntype of e = %T",e)

	f := 0.2
	fmt.Printf("\ntype of f = %T",f)

	fmt.Println("\ngA=",gA,"\ngB=",gB)

	// 声明多个变量
	var xx,yy int = 100,200
	fmt.Println("\nxx=",xx,"\nyy=",yy)

	var kk,ll = 100,"str"
	fmt.Println("\nkk=",kk,"\nll=",ll)
	
	// 多行多变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("\nvv=",vv,"\njj=",jj)
}