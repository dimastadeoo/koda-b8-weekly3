package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString(input string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}