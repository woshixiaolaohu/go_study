package base

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
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
