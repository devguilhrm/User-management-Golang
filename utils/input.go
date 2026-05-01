package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputStr(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInputInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// Printf with One Format
func ReadInputStr_oneF(prompt, format string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt, format)
	formatWithFunction, _ := reader.ReadString('\n')
	formatWithFunction = strings.TrimSpace(formatWithFunction)
	return formatWithFunction
}
