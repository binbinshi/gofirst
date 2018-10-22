package main

import (
	"fmt"
	"math"
	"strconv"
)

//可以有自定义行为
/**
	可以自己定义类型信息
	可以自定义 KB MB
 */
type(
	byte int8
	文本 string
	ByteSize int64
)

var(
	//常规方式
	ac = "hello"
	//使用并行方式进行类型推断
	bd, ce =1, 2
	// d := 3 不可以使用这种方式 :就是代替var的简写形式
)
func main(){
	var a int
	var b int32
	var c float64
	var d float32
	var e string
	var f bool
	//数组的高层封装slice
	var g []int
	var h [1]bool
	var j [1]int
	var k [1]byte

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt32)
	var l 文本
	l = "中文类型名"
	fmt.Println(l)

	/**
		可以设置类型别名

		go语言中类型安全，不存在隐似转换，类型推断
		变量的赋值
		变量名称 = 变量值
		var q int //变量的声明 全局变量的声明方式
		q = 1 //变量的赋值
		var 变量名称 变量类型 = 表达式
		var q int = 1
		类型推断
		var q = 1
		全局变量不能使用：=赋值
		q := 1
		方法体内变量的声明与赋值
		全局变量的声明可使用var()的方式进行简写

	*/

	q := 1
	fmt.Println(q)

	//多个变量的声明
	var d1, e1, f1 int
	//多个变量的赋值
	d1, e1, f1 = 2, 3, 4
	fmt.Println(d1,e1,f1)
	//多个变量同时声明和赋值
	var dd, ee, ff int = 5, 6, 7
	fmt.Println(dd,ee,ff)
	//省略变量类型，由系统推导
	var ii, jj, kk = 8, 9, 10
	fmt.Println(ii,jj,kk)
	//多个变量声明与赋值的最简单方法
	mm, nn, oo := 11, 12, 13
	fmt.Println(mm,nn,oo)
	//局部变量不可以使用var()的方式简写，只能使用并行方式
	fmt.Println(ac, bd, ce)
	// 空白符号 _  当有多个返回值时是有用的占位
	x, _, z := 1, 2, 3
	fmt.Println(x,z)
	// go 不存在隐式转换，所有类型必须显示声明
	//转换只能发生在两种相互兼容的类型之间
	var u float32 = 1.1
	v := int(u)
	fmt.Println(v, u)
	//类型转换的格式 <ValueA> [:] = <TypeOfValueA>(<ValueB>)
	/** 不可以转换 会报这个错误 cannot convert r (type bool) to type int
		var r bool = true
		t := int(r)
		fmt.Println(t)
	*/

	var vv int = 65
	rr := string(vv)
	fmt.Println(rr)

	tt := strconv.Itoa(vv)
	fmt.Println(tt)
	/*
		ss := strconv.Atoi(tt) 会如下报错 是因为函数返回的是多个返回值， 需要用到空宝字符
		multiple-value strconv.Atoi() in single-value context
	*/
	ss,  _ := strconv.Atoi(tt)
	fmt.Println(ss)

}