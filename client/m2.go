package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"os"
	"unsafe"
)

// ReadNdPrivatePem1
/**
 *  @Description: 读取私钥文件
 *  @param path 私钥文件路径
 *  @param pwd 私钥文件密码，无则填nil
 *  @return privateKey 私钥
 *  @return err
 */
func ReadNdPrivatePem1(path string) (privateKey *sm2.PrivateKey, err error) {
	// 打开文件读取私钥
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, fileInfo.Size(), fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	// 将pem格式私钥文件进行反序列化
	privateKey, err = x509.ReadPrivateKeyFromPem(buf, nil)
	if err != nil {
		return nil, err
	}
	return
}

// NdDecode1
/**
 *  @Description: SM2解密（私钥解密）
 *  @param cipherStr 加密后的字符串
 *  @param privateKey 私钥
 *  @return data 解密后的数据
 *  @return err
 */
func NdDecode1(cipherStr string, privateKey *sm2.PrivateKey) (data string, err error) {
	// 16进制字符串转[]byte
	bytes, _ := hex.DecodeString(cipherStr)
	// sm2解密
	var dataByte []byte
	dataByte, err = privateKey.DecryptAsn1(bytes)
	if err != nil {
		return data, err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&dataByte))
	return *str, err
}

func main() {
	s1:=make([]int, 0)
	s1 = append(s1, 2)
	s1 = append(s1, 3)

	s2:=s1
	s3:=make([]int, 0)
	fmt.Printf("%p, %p, %p\n", s1, s2, s3)
	fmt.Println(s1, s2, s3)
	s3 = s1

	fmt.Println(s1, s2, s3)
	fmt.Printf("%p, %p, %p\n", s1, s2, s3)
	//1、查询私钥文件
	//私钥
	//sign := "MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgvcMW3vYIUdE8kD29nDzqyXkkgDvzw5U8rokoBRjtjuOgCgYIKoEcz1UBgi2hRANCAASCDxuqrcwN6m0CpCzqB38ttaOMyG3b4D9pPI+DbmOH/Dw45wnWJg4+XuqCP8f5jdSnwecaIp2NLFvGW37TALBK"
	ndkjPrivateKey, err := ReadNdPrivatePem1("./client/ndkjPrivate.Pem")
	if err != nil {
		fmt.Println("读取私钥文件失败！")
	} else {
		fmt.Println("ndkj:", ndkjPrivateKey.PublicKey)
	}
	//2、密文
	cipherText := "04EE98A3AC2AEB53BAA1A6E1F0420A95D199594DF5D8DDD97E83E46A0E8EF0A99A2ECF40455E05171083F8D6DCC702AECD3FB810121809A111E4EEB365098E5F018BE7DC3684043019DA6D5B3F8DA3BAF3DAF9A562B26C435BC166ED10777456F1CB32F108FBAD"
	// data, _ := base64.StdEncoding.DecodeString(cipherText)
	//3、解密
	plainText, e := NdDecode1(cipherText, ndkjPrivateKey)
	if e != nil {
		fmt.Println("使用私钥解密失败！")
	}

	d2, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		fmt.Println("base64 err", err)
	}

	plaintxt, err:=  ndkjPrivateKey.DecryptAsn1([]byte(d2)) //sm2解密
	if err != nil {
		fmt.Println("使用私钥解密失败！")
	}
	//7、私钥解密内容
	fmt.Println("私钥解密内容：", plaintxt, plainText)

	d3, err := sm2.Decrypt(ndkjPrivateKey, []byte(cipherText), sm2.C1C2C3)
	if err != nil {
		fmt.Println("1111 使用私钥解密失败！", err)
	}
	fmt.Println("d3:", string(d3))
}
