package item

import "fmt"

func AddItem(name string, price float64) *Item {
	id := getNewItemId()
	i := &Item{
		ID:    id,
		Name:  name,
		Price: price,
	}

	if itemsList == nil {
		itemsList = make(map[int64]*Item)
	}

	itemsList[id] = i

	return i
}

var itemId int64 = 0

var itemsList map[int64]*Item

func getNewItemId() int64 {
	itemId += 1
	return itemId
}

func ListItems() {
	fs := "\n------------------ Item List ------------------\n"
	fs += "-----------------------------------------------\n"
	fs += fmt.Sprintf("| %-5s | %-14s | %-5s | %-10s |\n", "S.No", "Item Name", "ID", "Price")
	fs += "-----------------------------------------------\n"

	itemCount := 0

	for _, item := range itemsList {
		itemCount += 1
		fs += fmt.Sprintf("| %-5d | %-15s| %-5d | ₹%-9.2f |\n",
			itemCount, item.Name, item.ID, item.Price)
	}

	fs += "-----------------------------------------------\n"
	fs += fmt.Sprintf("| %-28s Total Items: %v |\n", "", itemCount)
	fs += "-----------------------------------------------\n"

	fmt.Println(fs)
}

func (i *Item) PrintItem() {
	fs := "\n------------| Item Details |-----------\n"
	fs += "---------------------------------------\n"
	fs += fmt.Sprintf("| %-5s | %-14s | %-10s |\n", "ID", "Item Name", "Price")
	fs += "---------------------------------------\n"

	fs += fmt.Sprintf("| %-5d | %-14s | ₹%-9.2f |\n",
		i.ID, i.Name, i.Price)

	fs += "---------------------------------------\n"

	fmt.Println(fs)
}

func FindItem(id int64) *Item {
	b := itemsList[id]

	return b
}

func DeleteItem(i *Item) Item {
	delete(itemsList, i.ID)
	return *i
}
