package main

import "fmt"

func testmap() {
	type Person struct {
		name string
		sex  string
		age  int
	}

	m := map[uint]Person{
		0: Person{"张无忌", "男", 18},
		1: Person{"周芷若", "女", 17},
	}

	//m[0].age += 1
	//整体更新结构体
	temp := m[0]
	temp.age += 1
	fmt.Println("age:", m[0].age)
	m[0] = temp
	fmt.Println(m)

}

func testmap2() {
	type Person struct {
		name string
		sex  string
		age  int
		good bool
	}

	people := map[uint]Person{
		0: Person{"张无忌", "男", 18, false},
		1: Person{"周芷若", "女", 17, false},
	}

	for name, _ := range people {
		// 必须改成指针才能修改
		if people[name].age < 50 {
			people[name].good = true
		}
	}

	m := map[int]string{
		0: "hzh",
		2: "5555",
	}
	for i, _ := range m {
		m[i] = "hhh"
	}
}

// map的key类型不能为slice、function、map
func main() {
	testmap2()
	testmap()
	type A struct {
		b int
		c string
	}
	s := make([]A, 0)
	s = append(s, A{})
	s = append(s, A{})
	s = append(s, A{})
	s = append(s, A{})
	d := make(map[*A]int)
	d1 := make(map[*A]int)
	i := 0
	for _, v := range s {
		i++
		d[&v] = i
	}
	for _, v := range s {
		i++
		v1 := v
		d1[&v1] = i
	}
	println(len(d), len(d1))
}
