package utils

import (
	perror "github.com/pkg/errors"
	_ "hello_grpc/test/link"
	"os"
)

func readFromFile(fp *os.File) string {
	return "dsdfd"
}
func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return perror.Errorf("open failed")
		//return perror.Wrap(err, "open failed")
	}
	defer f.Close()

	content := readFromFile(f)
	if len(content) == 0 {
		return perror.Errorf("content empty")
	}

	return nil
}

func helloWorld() string

func LinkHello() string {
	return helloWorld()
}
