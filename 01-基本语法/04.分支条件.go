package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func getSocreReult(score int) string {
	result := ""
	switch {
	//switch后可以没有表达式，默认自动加了brea，否则需要fallthrough
	case score < 0 || score > 100:
		result = "数字范围错误"
	case score < 60:
		result = "F"
	case score < 80:
		result = "C"
	case score < 90:
		result = "B"
	case score < 100:
		result = "C"
	case score == 100:
		result = "A+"
	}
	return result
}

func main() {
	const fileName = "scores.txt"
	_, err := ioutil.ReadFile(fileName)
	if (err != nil) {
		fmt.Println(err)
	} else {
		file, _ := os.Open(fileName)
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			score, err := strconv.Atoi(fileScanner.Text())
			result := ""
			if (err == nil) {
				result = strconv.Itoa(score) + " -> " + getSocreReult(score)
			} else {
				result = "非法数字 : " + fileScanner.Text()
			}
			fmt.Println(result)
		}
	}
}
