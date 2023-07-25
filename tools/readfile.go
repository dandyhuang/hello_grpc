package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Write0(fileName, buf string) {
	// strTest := "测试测试"

	var f *os.File
	var err error

	if CheckFileExist(fileName) { //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		if err != nil {
			fmt.Println("file open fail", err)
			return
		}
	} else { //文件不存在
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			fmt.Println("file create fail")
			return
		}
	}
	defer f.Close()
	//将文件写进去
	n, err1 := io.WriteString(f, buf)
	if err1 != nil {
		fmt.Println("write error", err1)
		return
	}
	fmt.Println("写入的字节数是：", n)
}

func main() {
	Write0("", "11\n")
	filename := flag.String("f", "", "Specify the file name")
	flag.Parse()

	if filename == nil || len(*filename) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*filename)
	if err != nil {
		log.Printf("Failed to open file %s,err:%s", *filename, err.Error())
		os.Exit(1)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	i := 0
	for {
		line, err := buf.ReadBytes('\n')
		log.Print(string(line))
		if i%2 == 0 {
			Write0("1.txt", string(line))
		} else {
			Write0("2.txt", string(line))
		}
		i++
		if err != nil {
			if err == io.EOF {
				log.Println("end of file")
				break
			} else {
				log.Printf("read file err:%s", err.Error())
				break
			}
		}
	}
}
