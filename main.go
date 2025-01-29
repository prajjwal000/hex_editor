package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "test", "filename")

	file, _ := os.Open(fileName)
	data := make([]byte, 100)
	file.Read(data)
	print(data)
}

func print(data []byte) error {
	var err error = nil

	fmt.Print("\t1\t2\t3\t4\t5\t6\t7\t8")
	for index, ch := range data {
		if index % 8 == 0 {
			fmt.Printf("\n%d\t",index/8)
		}
		fmt.Printf("%x\t",ch)
	}
	return err
}
