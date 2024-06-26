package base

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "base")
	// map声明
	var m1 map[string]string
	m2 := make(map[string]string)
	m2["name"] = "jelly"
	age, ok := m2["age"]
	if ok {
		println(age)
	} else {
		fmt.Println("m2中无age")
	}
	// 获取键值对
	m3 := make(map[string]int)
	m3["zhang_san"] = 18
	m3["lisi"] = 20
	m3["wang_wu"] = 21
	delete(m3, "zhang_san")
	for k, v := range m3 {
		fmt.Printf("%s's age is %d\n", k, v)
	}
	fmt.Println(m1 == nil, m2 == nil, m2, m2["name"])
	// slice声明
	var s1 []int
	s2 := make([]int, 10)
	s2[0] = 1
	s3 := s2[:5]
	fmt.Println(s1 == nil, s2 == nil, s2, cap(s3), len(s3))
	mapSort()
	// 生成map类型的切片
	s4 := make([]map[string]int, 2, 10)
	s4[0] = make(map[string]int)
	s4[1] = make(map[string]int)
	s4[0]["jelly"] = 26
	s4[0]["zhang_san"] = 20
	s4[1]["li_si"] = 19
	fmt.Printf("s类型:%T s的值：%v\n", s4, s4)
	// 生成切片类型的map
	m4 := make(map[string][]string, 1)
	m4["北京市"] = []string{"海淀区", "东城区", "西城区"}
	fmt.Printf("m类型：%T m的值：%v\n", m4, m4)
	// rune类型
	str1 := "大白兔"
	// 字符串转成rune切片
	str2 := []rune(str1)
	// 使用单引号修改字符 type为int32
	str2[1] = '黑'
	// 将rune类型str2转换为string
	fmt.Println(str1, str2, string(str2))
	// pointer指针
	n := 18
	p1 := &n
	v1 := *p1
	fmt.Printf("n的内存地址为%p,根据n内存地址取值为%d\v\n", p1, v1)
	// 99乘法表
	multiplicationTest()
	// 结构体
	type user struct {
		name       string
		age        int
		email      string
		privileged bool
	}
	// 使用机构类型声明变量并初始化为零值
	var bill user
	fmt.Printf("bill type is %T, value is %v\n", bill, bill)
	//lisa := user{
	//	name:       "lisa",
	//	age:        18,
	//	email:      "lisa@gmail.com",
	//	privileged: false,
	//}
	lisa := user{
		"lisa", 18, "lisa@gmail.com", false,
	}
	fmt.Printf("lisa type is %T, value is %v\n", lisa, lisa)
	type admin struct {
		person user
		level  string
	}
	jelly := admin{
		user{
			"jelly",
			26,
			"jelly@gmail.com",
			true,
		},
		"0",
	}
	fmt.Printf("jelly type is %T, value is %v\n", jelly, jelly)

	// 基于int64声明一个新类型
	//type Duration int64
	//var dur Duration
	//dur = int64(1000) //Cannot use 'int64(1000)' (type int64) as the type Duration
	n1 := Now()
	fmt.Printf("n1: %v\n", n1)
}

type Time struct {
	sec  int64
	nsec int32
	loc  *time.Location
}

type Duration int64

func Now() Time {
	sec, nsec := int64(1), int32(2)
	fmt.Printf("sec type is %T\n", sec)
	return Time{sec, nsec, time.Local}
}
func (t Time) Add(d Duration) Time {
	t.sec += int64(d / 10)
	nsec := int32(t.nsec) + int32(d%10)
	if nsec > 10 {
		t.sec++
		nsec -= 10
	} else if nsec < 0 {
		t.sec--
		nsec += 10
	}
	t.nsec = nsec
	return t
}

// 指针共享

// File  表示一个打开的文件描述符
type File struct {
	*file
}

// file 是*file的实际表示
// 额外的一层结构保证没有哪一个os的客户端能够覆盖这些数据，如果覆盖这些数据，可能在变量终结时关闭错误的文件描述符
type file struct {
	fd   int
	name string
	//dirinfo *dirInfo
	nepipe int32
}

func mapSort() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 声明初始化map
	sourceMap := make(map[string]int, 100)
	// 填充数据
	for i := 0; i < 100; i++ {
		// 生成stu开头的字符串
		key := fmt.Sprintf("stu%02d", 100-i)
		// 生成0-99随机整数
		value := rand.Intn(100)
		sourceMap[key] = value
	}
	keys := make([]string, 0, 100)
	for k, _ := range sourceMap {
		keys = append(keys, k)
	}
	// 对keys排序
	sort.Strings(keys)
	fmt.Printf(" keys.len: %d, keys.cap: %d\n", len(keys), cap(keys))
	//按照排序之后的keys遍历map
	//for _, key := range keys {
	//	fmt.Printf("key为%s的value为%d\n", key, sourceMap[key])
	//}
}

// 9*9乘法表
func multiplicationTest() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, i*j)
		}
		fmt.Println()
	}
}
