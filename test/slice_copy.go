package main

import "fmt"

func learnArray() {
	intarr := [5]int{12, 34, 55, 66, 43}
	slice := intarr[:]
	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
	fmt.Printf("address of slice 0x%p add of Arr 0x%p , %p\n", &slice[0], &intarr, &slice)
}
func learnSlice() {
	intarr := []int{12, 34, 55, 66, 43}
	slice := intarr[1:3]
	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
	fmt.Printf("address of slice 0x%p add of Arr 0x%p , %p\n", &slice[1], &intarr[1], &slice)
}

func main() {
	learnArray()
	learnSlice()
	// 设置元素数量为1000
	const elementCount = 10
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData[:]
	//addData := []int{100, 100, 100, 100, 100}
	//srcData = append(srcData, addData...)
	//fmt.Println("")
	fmt.Printf("%p %p %p %p\n", &refData[0], &srcData[0], &refData, &srcData)

	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	fmt.Println(copyData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}

	//使用内置的 copy 函数，复制切片
	var sliceSrc = []string{"Hello", "HaiCoder"}
	var sliceDst = []string{"嗨客网", "Python", "Golang"}
	copy(sliceDst, sliceSrc)
	fmt.Println("sliceDst =", sliceDst)
	sliceSrc[0] = "sdfsd"

	fmt.Println("sliceDst =", sliceDst)
}
