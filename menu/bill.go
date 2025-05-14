package menu

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/onebitgod/billsystem/bill"
)

func BillMenu(clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	fmt.Println("---------- Bill Menu ----------")
	fmt.Println("\n| \033[33mv\033[0m: View Bills | \033[33mn\033[0m: New Bill | \033[33mm\033[0m: Main Menu |")
	fmt.Print("\nChoose your option : ")

	for {
		input := ReadFromUser(reader)

		switch input {
		case "v":
			ViewBill(true)
			return
		case "n":
			NewBill(true)
			return
		case "m":
			MainMenu(true)
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mUnknown Command\033[0m", input)
			BillMenu(false)
			return
		}
	}
}

func ViewBill(clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	bill.ListBills()

	fmt.Println("\n\033[33mChoose Below option or Enter Bill Id to proceed\033[0m")

	fmt.Println("\n| \033[33md\033[0m: Delete Bill | \033[33mb\033[0m: Bill Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit |")

	fmt.Print("\nChoose your option : ")

	for {
		input := ReadFromUser(reader)

		switch input {
		case "d":
			DeleteBill(true)
			return
		case "b":
			BillMenu(true)
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
				ViewBill(false)
			} else {
				b := bill.FindBill(id)
				if b == nil {
					fmt.Print("\033[H\033[2J")
					fmt.Println("\033[31mBill Not Found with ID\033[0m", id)
					ViewBill(false)
					return
				}
				EditBill(b)
			}
			return
		}
	}
}

func DeleteBill(clear bool) {

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter Bill Id To be deleted : ")

	input := ReadFromUser(reader)

	if id, err := strconv.ParseInt(input, 10, 64); err != nil {
		fmt.Print("\033[H\033[2J")
		fmt.Println("\033[31mUnknown Command\033[0m", input)
		DeleteBill(false)
		return
	} else {
		b := bill.FindBill(id)
		if b == nil {
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mBill Not Found with ID\033[0m", id)
			DeleteBill(false)
			return
		}
		b.PrintBill()
		fmt.Print("\nAre you sure want to delete above bill?")
		fmt.Println("\n| \033[33my\033[0m: Yes | \033[33mn\033[0m: No | \033[33mb\033[0m: Bill Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit |")
		fmt.Print("\nChoose your option : ")

		input := ReadFromUser(reader)

		switch input {
		case "y":
			bill.DeleteBill(b)
			ViewBill(true)
			return
		case "n":
			ViewBill(true)
			return
		case "b":
			BillMenu(true)
			return
		case "m":
			MainMenu(true)
			return
		case "q":
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mUnknown Command\033[0m", input)
			ViewBill(true)
			return
		}
	}

	// Execute something as soon as input is entered

}

func EditBill(b *bill.Bill) {

	if b == nil {
		if b = SearchBill(true); b == nil {
			return
		}
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\033[H\033[2J")

	b.PrintBill()

	fmt.Println("\n| \033[33ma\033[0m: Add Item | \033[33md\033[0m: Delete Item | \033[33me\033[0m: Edit Customer | \033[33mb\033[0m: Bill Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit |")

	fmt.Print("\nChoose your option : ")

	for {
		input := ReadFromUser(reader)

		// Execute something as soon as input is entered
		switch input {
		case "a":
			AddItemToBill(b, true)
			return
		case "d":
			DeleteItemFromBill(b, true)
			return
		case "e":
			editCustomerName(b, true)
			return
		case "b":
			BillMenu(true)
			return
		case "m":
			MainMenu(true)
			return
		case "q":
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("\033[31mUnknown Command\033[0m", input)
			ViewBill(false)
			return
		}
	}
}

func SearchBill(clear bool) *bill.Bill {
	reader := bufio.NewReader(os.Stdin)

	var b *bill.Bill = nil

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	fmt.Println("---------- Search Bill ----------")

	fmt.Println("| \033[33mb\033[0m: Bill Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit|")

	fmt.Print("\nEnter the Bill ID : ")

	input := ReadFromUser(reader)
	// Execute something as soon as input is entered
	switch input {
	case "b":
		BillMenu(true)
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
			b = SearchBill(false)
		} else {
			b = bill.FindBill(id)
			if b == nil {
				fmt.Print("\033[H\033[2J")
				fmt.Println("\033[31mBill Not Found with ID\033[0m", id)
				b = SearchBill(false)
			}

		}
		return b

	}
}

func AddItemToBill(b *bill.Bill, clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	item := SearchItem(true)

	if item == nil {
		return
	}

	fmt.Println("\nItem to be Added")
	item.PrintItem()

	fmt.Println("\n| \033[33mc\033[0m: Change Item | \033[33mb\033[0m: Bill Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit|")
	fmt.Print("\nEnter the Item Quantity : ")

	input := ReadFromUser(reader)

	switch input {
	case "c":
		AddItemToBill(b, true)
	case "b":
		BillMenu(true)
	case "m":
		MainMenu(true)
		return
	case "q":
		return
	default:
		if qty, err := strconv.Atoi(input); err != nil {
			fmt.Println("\033[31mInvalid Input\033[0m", input)
			EditBill(b)
		} else {
			b.AddItem(*item, qty)
			EditBill(b)
		}

		return

	}
}

func DeleteItemFromBill(b *bill.Bill, clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	item := SearchItem(true)

	if item == nil {
		return
	}

	fmt.Println("\nItem to be Deleted")
	item.PrintItem()

	fmt.Println("\n| \033[33mp\033[0m: Proceed | \033[33mc\033[0m: Change Item | \033[33mb\033[0m: Bill Menu | \033[33mm\033[0m: Main Menu | \033[33mq\033[0m: exit|")
	fmt.Print("\nChoose your option : ")

	input := ReadFromUser(reader)

	switch input {
	case "c":
		DeleteItemFromBill(b, true)
	case "b":
		BillMenu(true)
	case "m":
		MainMenu(true)
		return
	case "p":
		b.DeleteItem(*item)
		EditBill(b)
		return
	case "q":
		return

	default:
		fmt.Print("\033[H\033[2J")
		fmt.Println("\033[31mInvalid Input\033[0m", input)
		DeleteItemFromBill(b, false)
		return

	}
}

func NewBill(clear bool) {

	b := bill.CreateBill("")

	editCustomerName(b, clear)

}

func editCustomerName(b *bill.Bill, clear bool) {
	reader := bufio.NewReader(os.Stdin)

	if clear {
		fmt.Print("\033[H\033[2J")
	}

	b.PrintBill()

	isValidName := false

	name := ""
	for !isValidName {

		fmt.Printf("\nEnter Customer's New Name (%v) : ", b.CustomerName)
		name = ReadFromUser(reader)
		re := regexp.MustCompile(`^[A-Za-z]{3,}$`)
		if isValidName = re.MatchString(name); !isValidName {
			fmt.Print("\033[H\033[2J")
			b.PrintBill()
			fmt.Println("\033[31mInvalid Name\033[0m", name)

		} else {
			b.CustomerName = name
			EditBill(b)

		}

	}

}
