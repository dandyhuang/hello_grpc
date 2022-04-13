package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

type OrderedMap struct {
	Order []string
	Map   map[string]string
}

func (om *OrderedMap) UnmarshalJSON(b []byte) error {
	json.Unmarshal(b, &om.Map)

	index := make(map[string]int)
	for key := range om.Map {
		om.Order = append(om.Order, key)
		esc, _ := json.Marshal(key) //Escape the key
		index[key] = bytes.Index(b, esc)
	}

	sort.Slice(om.Order, func(i, j int) bool { return index[om.Order[i]] < index[om.Order[j]] })
	return nil
}

func (om OrderedMap) MarshalJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	buf.WriteRune('{')
	l := len(om.Order)
	for i, key := range om.Order {
		km, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(km)
		buf.WriteRune(':')
		vm, err := json.Marshal(om.Map[key])
		if err != nil {
			return nil, err
		}
		fmt.Println("vm:", string(vm))
		buf.Write(vm)
		if i != l-1 {
			buf.WriteRune(',')
		}
		fmt.Println(buf.String())
	}
	buf.WriteRune('}')
	fmt.Println(buf.String())
	return buf.Bytes(), nil
}

type Student struct {
	sex  *int
	name string
}

type person struct {
	Name string
}

func main() {
	//var s1 Student
	//var sex int
	//sex = 2
	//s1.sex = &sex
	//s1.name = "hahh"
	//p := (*[2]unsafe.Pointer)(unsafe.Pointer(&s1))[1]
	//pi := (*Student)(p)
	//fmt.Println(pi, &s1)
	p := person{}
	var o map[string]interface{}

	src := "{\"name\":\"你好的师傅打了\"}"
	json.Unmarshal([]byte(src), &p)

	fmt.Println(p)

	obj := `{"key3":"value3","key2":"value2","key1":"value1"}`
	json.Unmarshal([]byte(obj), &o)
	fmt.Println(o)
	r, _ := json.Marshal(o)
	fmt.Println(string(r))
	var o1 OrderedMap
	o1.UnmarshalJSON([]byte(obj))
	fmt.Println("o1:", o1)
	s, _ := o1.MarshalJSON()
	fmt.Println("finale:", string(s))

}
