package menu

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadFromUser(reader *bufio.Reader) string {

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	input = strings.TrimSpace(input)

	return input

}
