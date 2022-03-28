package main

import "fmt"

func main() {
	//var c byte
	//var ans []byte
	//m := map[byte]bool{}
	//for {
	//	n, _ := fmt.Scanf("%c", &c) //此处不能用fmt.Scan()来读取
	//	if n == 0 {
	//		break
	//	}
	//	if m[c] == false {
	//		m[c] = true
	//		ans = append(ans, c)
	//	}
	//}
	//fmt.Println(string(ans))
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t\t", name, age, married)

}
