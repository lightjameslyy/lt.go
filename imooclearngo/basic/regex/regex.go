package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is lightjameslyy@gmail.com
another email is lightjames@163.com
another   :    lightjames@edu.ac.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	// match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
