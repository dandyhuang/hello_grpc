package main

import (
	"fmt"
	"reflect"
)

type MQOption struct {
	TopicName string `json:"topic_name" comment:"消息队列名称"`
}

func getStructTag(v interface{}, tag string) {
	structVal := reflect.ValueOf(v).Elem()
	structType := structVal.Type()
	fmt.Println(structType, structVal)
	fmt.Println("size:", structVal.NumField())
	for i := 0; i < structVal.NumField(); i++ {
		fieldType := structType.Field(i)
		fmt.Println("fieldType", fieldType, "name:", fieldType.Name)
		field := structVal.FieldByName(fieldType.Name)
		fmt.Println("field", field)

		comment := fieldType.Tag.Get(tag)
		fmt.Println("tag", comment)
	}
	structType1 := reflect.TypeOf(v).Elem()
	// 遍历结构体所有成员
	for i := 0; i < structType1.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := structType1.Field(i)
		// 输出成员名和tag
		fmt.Printf("type1 name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}

	// 通过字段名, 找到字段类型信息
	if cType, ok := structType1.FieldByName("TopicName"); ok {
		// 从tag中取出需要的tag
		fmt.Println("topicName:", cType.Tag.Get("json"), cType.Tag.Get("comment"))
	}
}
func getReflectType() {
	type student struct {
	}
	s := &student{}
	tfs := reflect.TypeOf(s)
	// 显示反射类型对象的名称和种类和元素
	fmt.Println("type:", tfs.Name(), tfs.Kind(), tfs.Elem())
	// 取类型的元素 It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice
	elems := tfs.Elem()
	fmt.Println("elemstype:", elems.Name(), elems.Kind())
	s1 := student{}
	tfs1 := reflect.TypeOf(s1)
	fmt.Println("type1:", tfs1.Name(), tfs1.Kind())
}

func main() {
	mo := new(MQOption)
	getStructTag(mo, "comment")
	// 反射的类型（Type）与种类（Kind）
	getReflectType()
}
