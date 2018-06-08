package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// func ReadLine(filename string) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	r := bufio.NewReaderSize(f, 4*1024)
// 	line, isPrefix, err := r.ReadLine()
// 	for err == nil && !isPrefix {
// 		s := string(line)
// 		fmt.Println(s)
// 		line, isPrefix, err = r.ReadLine()
// 	}
// 	if isPrefix {
// 		fmt.Println("buffer size to small")
// 		return
// 	}
// 	if err != io.EOF {
// 		fmt.Println(err)
// 		return
// 	}
// }

func ReadString(filename string) string {
	var result = ""

	f, err := os.Open(filename)
	if err != nil {
		// fmt.Println(err)
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	result += line
	for err == nil {
		// fmt.Print(line)
		line, err = r.ReadString('\n')
		result += line
	}
	if err != io.EOF {
		// fmt.Println(err)
		return ""
	}
	return result
}

func ReadLines(filename string) []string {
	var txt = ReadString(filename)
	var lines = strings.Split(txt, "\n")
	return lines
}

func ReadCSV(filename string) [][]string {
	var lines = ReadLines(filename)
	var result = make([][]string, len(lines))
	for index := 0; index < len(lines); index++ {
		result[index] = strings.Split(lines[index], ",")
	}
	return result
}

func ReadSetting(filename string) map[string]string {
	var lines = ReadLines(filename)

	var result = make(map[string]string)
	for index := 0; index < len(lines); index++ {
		var splits = strings.Split(lines[index], "=")
		result[splits[0]] = splits[1]
	}
	return result
}
