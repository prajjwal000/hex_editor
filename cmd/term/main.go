package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fileRead string
	var fileWrite string
	var writeBool bool
	flag.StringVar(&fileRead, "readFile", "cmd/term/testread", "filename")
	flag.StringVar(&fileWrite, "writeFile", "cmd/term/testwrite", "filename")
	flag.BoolVar(&writeBool, "w", false, "Enable to write")
	flag.Parse()

	writeFile, _:= os.OpenFile(fileWrite, os.O_RDWR|os.O_CREATE, 0644)
	readFile, err := os.Open(fileRead)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()
	defer writeFile.Close()

	if writeBool == false {
		fi, _ := readFile.Stat()
		data := make([]byte, fi.Size())
		readFile.Read(data)
		print(data)
	} else {
		write(writeFile, readFile)
	}
}

func print(data []byte) error {
	var err error = nil

	// fmt.Print("\t1\t2\t3\t4\t5\t6\t7\t8")
	for index, ch := range data {
		if index%8 == 0 && index != 0 {
			helperString := string(data[index-8 : index])
			helperString = strings.ReplaceAll(helperString, "\n", "")
			helperString = strings.ReplaceAll(helperString, "\t", "")
			fmt.Println(" | ", helperString)
			fmt.Printf("\n")
		}
		fmt.Printf("%x\t", ch)
	}
	return err
}

func write(writeFile, readFile *os.File) error {
	var err error = nil
	fi, _ := readFile.Stat()
	data := make([]byte, fi.Size())
	var towrite []string
	readFile.Read(data)

	byteVal := ""
	status := true
	for _, ch := range data {
		if ch == '|' {
			status = false
		}
		if isHexString(ch) && status{
			byteVal += string(ch)
		}
		if ch == '\n' || ch == '\t' && status {
			towrite = append(towrite, byteVal)
			byteVal = ""
			status = true
		}
	}

	var towritefinal []byte
	for i := range towrite {
		hexByte, _ := strconv.ParseInt(towrite[i], 16, 0)
		towritefinal = append(towritefinal, byte(hexByte))
	}

	writeFile.Write(towritefinal)

	return err
}

//Helper functions

func isHexString(ch byte) bool {
	return ('a' <= ch && ch <= 'f') || ('0' <= ch && ch <= '9')
}
