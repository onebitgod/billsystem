package main

import (
	"github.com/onebitgod/billsystem/bill"
	"github.com/onebitgod/billsystem/item"
	menu "github.com/onebitgod/billsystem/menu"
)

func main() {

	i1 := item.AddItem("Pasta", 100)
	i2 := item.AddItem("Pizza", 150)
	// item.ListItems()

	b1 := bill.CreateBill("Himanshu")

	b1.AddItem(*i1, 2)
	b1.AddItem(*i2, 3)

	// i1.PrintItem()

	// b1.PrintBill()

	// bill.ListBills()

	menu.MainMenu(true)

}
