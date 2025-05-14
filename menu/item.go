package menu

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/onebitgod/billsystem/item"
)

func ItemMenu(clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	fmt.Println("---------- Item Menu ----------")
	fmt.Println("\n| \033[33mv\033[0m: View Items | \033[33ma\033[0m: Add Item | \033[33mm\033[0m: Main Menu |")
	fmt.Print("\nChoose your option : ")

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		// Execute something as soon as input is entered
		switch input {
		case "v":
			ViewItem(true)
			return
		case "a":
			NewItem(true)
			return
		case "m":
			MainMenu(true)
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mUnknown Command\033[0m", input)
			ItemMenu(false)
			return
		}
	}
}

func SearchItem(clear bool) *item.Item {
	reader := bufio.NewReader(os.Stdin)

	var i *item.Item = nil

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	item.ListItems()

	fmt.Println("---------- Search Item ----------")

	fmt.Println("| \033[33mb\033[0m: Bill Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit|")

	fmt.Print("\nEnter the Item ID : ")

	for {
		input := ReadFromUser(reader)
		// Execute something as soon as input is entered
		switch input {
		case "i":
			ItemMenu(true)
			return nil
		case "m":
			MainMenu(true)
			return nil
		case "q":
			return nil
		default:
			if id, err := strconv.ParseInt(input, 10, 64); err != nil {
				fmt.Print("\033[H\033[2J")
				fmt.Println("\033[31mUnknown Command\033[0m", input)
				i = SearchItem(false)
			} else {
				i = item.FindItem(id)
				if i == nil {
					fmt.Print("\033[H\033[2J")
					fmt.Println("\033[31mItem Not Found with ID\033[0m", id)
					i = SearchItem(false)
				}

			}
			return i

		}
	}
}

func ViewItem(clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	item.ListItems()

	fmt.Println("\n\033[33mChoose Below option or Enter Item Id to proceed\033[0m")

	fmt.Println("\n| \033[33md\033[0m: Delete Item | \033[33mi\033[0m: Item Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit |")

	fmt.Print("\nChoose your option : ")

	for {
		input := ReadFromUser(reader)

		switch input {
		case "d":
			DeleteItem(true)
			return
		case "i":
			ItemMenu(true)
			return
		case "m":
			MainMenu(true)
			return
		case "q":
			return
		default:
			if id, err := strconv.ParseInt(input, 10, 64); err != nil {
				fmt.Print("\033[H\033[2J")
				fmt.Println("\033[31mUnknown Command\033[0m", input)
				ViewItem(false)
			} else {
				i := item.FindItem(id)
				if i == nil {
					fmt.Print("\033[H\033[2J")
					fmt.Println("\033[31mItem Not Found with ID\033[0m", id)
					ViewItem(false)
					return
				}
				EditItem(i)
			}
			return
		}
	}
}

func DeleteItem(clear bool) {

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter Item Id To be deleted : ")

	input := ReadFromUser(reader)

	if id, err := strconv.ParseInt(input, 10, 64); err != nil {
		fmt.Print("\033[H\033[2J")
		fmt.Println("\033[31mUnknown Command\033[0m", input)
		DeleteItem(false)
		return
	} else {
		b := item.FindItem(id)
		if b == nil {
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mItem Not Found with ID\033[0m", id)
			DeleteItem(false)
			return
		}
		b.PrintItem()
		fmt.Print("\nAre you sure want to delete above Item?")
		fmt.Println("\n| \033[33my\033[0m: Yes | \033[33mn\033[0m: No | \033[33mi\033[0m: Item Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit |")
		fmt.Print("\nChoose your option : ")

		input := ReadFromUser(reader)

		switch input {
		case "y":
			item.DeleteItem(b)
			ViewItem(true)
			return
		case "n":
			ViewItem(true)
			return
		case "b":
			ItemMenu(true)
			return
		case "m":
			MainMenu(true)
			return
		case "q":
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mUnknown Command\033[0m", input)
			ViewItem(true)
			return
		}
	}

	// Execute something as soon as input is entered

}

func EditItem(i *item.Item) {

	if i == nil {
		if i = SearchItem(true); i == nil {
			return
		}
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\033[H\033[2J")

	i.PrintItem()

	fmt.Println("\n| \033[33mp\033[0m: Edit Price | \033[33mn\033[0m: Edit Name | \033[33mi\033[0m: Item Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit |")

	fmt.Print("\nChoose your option : ")

	for {
		input := ReadFromUser(reader)

		switch input {
		case "p":
			editPrice(i, true)
			EditItem(i)
			return
		case "n":
			editName(i, true)
			EditItem(i)
			return
		case "i":
			ItemMenu(true)
			return
		case "m":
			MainMenu(true)
			return
		case "q":
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mUnknown Command\033[0m", input)
			ViewItem(false)
			return
		}
	}
}

func editPrice(i *item.Item, clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	i.PrintItem()

	fmt.Printf("\nEnter Item's New Price (₹%v) : ", i.Price)
	input := ReadFromUser(reader)

	if price, err := strconv.ParseFloat(input, 64); err == nil {
		i.Price = float64(price)

	} else {
		fmt.Print("\033[H\033[2J")
		fmt.Println("\033[31mInvalid Input\033[0m", input)
		editPrice(i, false)
	}

}

func NewItem(clear bool) {

	i := item.AddItem("", 0)

	editName(i, clear)
	editPrice(i, clear)
	EditItem(i)

}

func editName(i *item.Item, clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	i.PrintItem()

	isValidName := false

	name := ""
	for !isValidName {

		fmt.Printf("\nEnter Item's New Name (%v) : ", i.Name)
		name = ReadFromUser(reader)
		re := regexp.MustCompile(`^[A-Za-z]{3,}$`)
		if isValidName = re.MatchString(name); !isValidName {
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mInvalid Name\033[0m", name)
		} else {
			i.Name = name
		}

	}

}
