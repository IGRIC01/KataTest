package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	arr := []rune(t.Content)
	sum := 0

	for strings.Contains(string(arr), "  ") {
		arr = []rune(strings.ReplaceAll(string(arr), "  ", " "))
	}

	for index, value := range arr {
		if value == '-' {
			arr[index-1], arr[index+1] = arr[index+1], arr[index-1]
		}
	}

	arr = []rune(strings.ReplaceAll(string(arr), "-", ""))

	for index, value := range arr {
		if value == '+' {
			arr[index] = '!'
		}
	}

	for _, char := range arr {
		if unicode.IsDigit(char) {
			ch, _ := strconv.Atoi(string(char))
			sum += ch
			arr = []rune(strings.Replace(string(arr), string(char), "", 1))
		}
	}

	num := strconv.Itoa(sum)

	if num == "0" {
		fmt.Println(string(arr))
	} else {
		fmt.Println(string(arr) + " " + num)
	}

}

func main() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строчку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}

}
