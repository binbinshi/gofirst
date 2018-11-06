package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aa"
	fmt.Printf("equalfold aa compare %s = %t\n", s, strings.EqualFold(s, "aa"))
	var t = "AA"
	fmt.Printf("equalfold aa compare %s = %t\n", t, strings.EqualFold(t, "aa"))
	var n string = "bb"
	fmt.Printf("equalfold aa compare %s = %t\n", n, strings.EqualFold(n, "aa"))
	// 将单词中的首字母改为标题格式的字符串
	m := strings.Title(n)
	fmt.Printf("Title format %s =  %s\n", n, m)
	fmt.Printf("equalfold bb compare %s = %t\n", m, strings.EqualFold(m, "bb"))
	p := "waht is go go go"
	fmt.Printf("HasPrefix wa in %s = %t\n", p, strings.HasPrefix(p, "wa"))
	//判断是否包含子串 func Contains(s, substr string)
	fmt.Printf("Contains go in %s = %t\n ", p, strings.Contains(p, "go"))
	fmt.Printf("Contains  in %s = %t\n", p, strings.Contains(p, " "))
	//判断字符串中是否包含字符串chars中的任一字符
	fmt.Printf("ContainsAny wa in %s = %t\n", p, strings.ContainsAny(p, "wa"))
	//返回字符串中有几个不重复的子串
	fmt.Printf("Count go in %s = %d\n", p, strings.Count(p, "go"))

	//获取字符串位置的函数 必须显示的声明调用方法所在的包
	fmt.Printf("Index go in %s = %d\n", p, strings.Index(p, "go"))
	//获取字符在字符串中第一次出现的位置，不存在返回-1
	fmt.Printf("IndexByte  w in %s = %d\n", p, strings.IndexByte(p, 'w'))
	//字符串chars中的任一utf-8码值在s中第一次出现的位置，会匹配第一个关联的字符，不是总体的串出现的索引地址如果不存在或者chars为空字符串则返回-1
	fmt.Printf("IndexAny  g  in %s = %d\n", p, strings.IndexAny(p, " g"))
	//字符串中第一个满足函数f的位置i（该处的utf-8码值r满足f(r) == true, 不存在则返回-1）
	p1 := "C:\\Windows\\System32"
	fmt.Printf("IndexFunc in go %s = %d\n", p1, strings.IndexFunc(p1, isSlash))
	//子串在字符串中最后一次出现的位置，不存在时返回-1
	fmt.Printf("LastIndex go %s = %d\n", p, strings.LastIndex(p, "go"))

	//字符串中字符处理
	//将所有的字符串中的字符转为对应的小写版本的拷贝
	p2 := "AAAAA BBBBB"
	fmt.Printf("ToLower %s = %s\n", p2, strings.ToLower(p2))
	//将所有的字符串中的字母变成对应的大写版本的拷贝
	fmt.Printf("ToUpper %s = %s\n", p, strings.ToUpper(p))
	//返回n个传入字符串的拷贝
	fmt.Printf("Repeat %s * %d = %s\n", s, 3, strings.Repeat(s, 3))
	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	//func Replace (s, old, new stirng, n int)string
	fmt.Printf("Replace %s old : %s replaceto new : %s replace number %d  = %s\n", p, "go", "java", 2, strings.Replace(p, "go", "java", 2))
	fmt.Printf("Replace %s old : %s replaceto new : %s replace number %d  = %s\n", p, " ", "gm", -1, strings.Replace(p, " ", "gm", -1))
	//将字符串中的每一个unicode码值r都替换为mapping（r）,返回这些新码值组成的字符串拷贝。如果mapping返回一个负值，将会丢弃该码值而不会被替换
	//func Map(mapping func(rune) rune, s string)
	fmt.Printf("Map  %s = %s\n", p, strings.Map(func(r rune) rune {
		if r == 'g' {
			return 'b'
		}
		return r
	}, p))

	//字符串前后端的处理
	//返回字符串前后端所有cutset包含的utf-8码值都去掉的字符串
	p3 := "java ddd java"
	fmt.Printf("Trim %s = %s\n", p3, strings.Trim(p3, "java"))
	//返回将入参前后的所有空白都去掉的字符串
	fmt.Printf("TrimSpace %s = %s\n", p3, strings.TrimSpace("ddd"))
	//去除所有满足函数要求的字符串
	p4 := "jjjj sxxxxxx jjddjj"
	fmt.Printf("TrimFunc %s = %s\n", p4, strings.TrimFunc(p4, func(r rune) bool {
		if r == 'j' {
			return true
		}
		return false
	}))
	//将字符串按照空白分割的多个字符串
	fmt.Printf("Fields %s = %v\n", p, strings.Fields(p))
	//按照传入的字符串进行分割并将字符串去除，空格会被保留，
	p5 := "aasbbs ffsffggsss"
	p6 := strings.Split(p5, "s")
	fmt.Printf("Split %s = %v ： len %d\n", p5, p6, len(p6))
	//字符串拼接
	fmt.Printf("Join %v = %s9end len: %d \n", p6, strings.Join(p6, "1"), len(strings.Join(p6, "1")))
}

func isSlash(r rune) bool {
	return r == '\\' || r == '/'
}
