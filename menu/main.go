package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MainMenu(clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	fmt.Println("---------- Main Menu ----------")
	fmt.Println("\n| \033[33mb\033[0m: Bills Menu | \033[33mi\033[0m: Items Menu | \033[33mq\033[0m: Exit |")
	fmt.Print("\nChoose your option : ")

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		switch input {

		case "b":
			BillMenu(true)
			return

		case "i":
			ItemMenu(true)
			return

		case "q":
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[32mThank You\033[0m")
			return

		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mUnknown Command\033[0m", input)
			MainMenu(false)
			return
		}
	}
}
