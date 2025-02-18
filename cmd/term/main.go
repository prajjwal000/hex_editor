package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var writeBool bool
	var fileName string
	flag.BoolVar(&writeBool, "w", false, "Enable to write")
	flag.Parse()

	if writeBool {
		fileName = os.Args[2]
	} else {
		fileName = os.Args[1]
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if writeBool == false {
		fmt.Print(print(data))
	} else {
		tmpFile, err := os.Create("tmp")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpFile.Name())
		defer fmt.Print("Defer")

		tmpString, _ := print(data)
		tmpFile.WriteString(tmpString)

		cmd := exec.Command("nvim", tmpFile.Name())
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Waiting for editing\n")
		err = cmd.Wait()
		fmt.Print("Editing finished\n")
		fmt.Print("Writing hex\n")
		tmpFile.Sync()
		tmpFile.Close()

		tmpFile, err = os.Open("tmp")
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening ", fileName)
		}
		write(tmpFile, file)
		fmt.Print("Written\n")
	}
}

func print(data []byte) (string, error) {
	var err error = nil
	ret := strings.Builder{}
	// fmt.Print("\t1\t2\t3\t4\t5\t6\t7\t8")
	for index, ch := range data {
		if index%8 == 0 && index != 0 {
			helperString := string(data[index-8 : index])
			helperString = strings.ReplaceAll(helperString, "\n", "")
			helperString = strings.ReplaceAll(helperString, "\t", "")
			ret.WriteString(" | " + helperString)
			ret.WriteString(fmt.Sprintln())
		}
		ret.WriteString(fmt.Sprintf("%x\t", ch))
	}
	return ret.String(), err
}

func write(parseFile, writeFile *os.File) {
	l, _ := parseFile.Stat()
	data := make([]byte, l.Size())
	writebytes := make([]byte, l.Size())
	parseFile.Read(data)
	if string(data) == "" {
		log.Fatal("Zero bytes\n")
	}

	status := true
	helperString := ""
	i := 0
	for _, v := range data {
		fmt.Println(v, helperString, string(writebytes))
		if isHexString(v) && status {
			helperString += string(v)
		}
		if (v == '\t') && status {
			byes, err := strconv.ParseInt(helperString, 16, 0)
			if err != nil {
				log.Print(err, v)
			}
			writebytes[i] = byte(byes)
			helperString = ""
			i++
		}

		if v == '|' {
			status = false
		}

		if v == '\n' {
			status = true
		}
	}

	fmt.Print(writebytes)

	writeFile.Write(writebytes)

}

//Helper functions

func isHexString(ch byte) bool {
	return ('a' <= ch && ch <= 'f') || ('0' <= ch && ch <= '9')
}
