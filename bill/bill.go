package bill

import (
	"fmt"
	"time"

	"github.com/onebitgod/billsystem/item"
)

func CreateBill(customerName string) *Bill {
	id := getNewItemId()
	b := &Bill{
		ID:           id,
		CustomerName: customerName,
		Total:        0,
		Tip:          0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if billsList == nil {
		billsList = make(map[int64]*Bill)
	}

	billsList[id] = b

	return b
}

var billId int64 = 0

var billsList map[int64]*Bill

func getNewItemId() int64 {
	billId += 1
	return billId
}

func (b *Bill) UpdateTip(tip float64) Bill {
	b.Tip = tip
	return *b
}

func (b *Bill) AddItem(item item.Item, count int) Bill {

	if bi, _ := b.FindBillItem(item.ID); bi != nil {
		bi.Count += count
	} else {
		b.Items = append(b.Items, BillItem{
			Item:  item,
			Count: count,
		})
	}

	b.Total = b.Total + item.Price*float64(count)

	return *b
}

func (b *Bill) DeleteItem(item item.Item) Bill {

	if bi, i := b.FindBillItem(item.ID); bi != nil {
		b.Items = append(b.Items[:i], b.Items[i+1:]...)
		b.Total = b.Total - item.Price*float64(bi.Count)
	}

	return *b
}

func DeleteBill(b *Bill) Bill {
	delete(billsList, b.ID)
	return *b
}

func (b *Bill) FindBillItem(id int64) (*BillItem, int) {

	var bi *BillItem = nil
	var index int = 0

	for i := 0; i < len(b.Items); i++ {
		if b.Items[i].Item.ID == id {
			index = i
			bi = &b.Items[i]
		}
	}

	return bi, index

}

func (b *Bill) PrintBill() {
	fs := "\n---------------------------| Bill Details |-------------------------\n"
	fs += "--------------------------------------------------------------------\n"
	fs += fmt.Sprintf("|  Customer : %-44s | ID: %d |\n", b.CustomerName, b.ID)
	fs += "--------------------------------------------------------------------\n"
	fs += "--------------------------------------------------------------------\n"
	fs += fmt.Sprintf("| %-5s | %-14s | %-5s | %-5s | %-10s | %-10s |\n", "S.No", "Item Name", "ID", "Qty", "Price/Qty", "Total")
	fs += "--------------------------------------------------------------------\n"

	total := 0.0
	for i, item := range b.Items {
		subTotal := item.Item.Price * float64(item.Count)
		total += subTotal
		fs += fmt.Sprintf("| %-5d | %-15s| %-5d | %-5d | ₹%-10.2f | ₹%-8.2f |\n",
			i+1, item.Item.Name, item.Item.ID, item.Count, item.Item.Price, subTotal)
	}

	fs += "--------------------------------------------------------------------\n"
	fs += fmt.Sprintf("| %-49s Total: ₹%0.2f |\n", "", b.Total)
	fs += "--------------------------------------------------------------------\n"

	fmt.Println(fs)
}

func ListBills() {
	fs := "\n----------------------------------- Bills List ---------------------------------------\n"
	fs += "-------------------------------------------------------------------------------------\n"
	fs += fmt.Sprintf("| %-5s | %-14s | %-5s | %-5s | %-18s | %-10s |\n", "S.No", "Customer Name", "BillID", "Total items", "CreatedAt", "Total Amount")

	fs += "-------------------------------------------------------------------------------------\n"

	totalBills := 0

	for _, bill := range billsList {
		totalBills += 1
		fs += fmt.Sprintf("| %-5d | %-15s| %-6d | %-11d | %-10v | ₹%-12v|\n",
			totalBills, bill.CustomerName, bill.ID, len(bill.Items), bill.CreatedAt.Format("02 Jan 2006, 15:04"), bill.Total)
	}

	fs += "-------------------------------------------------------------------------------------\n"
	fs += fmt.Sprintf("| %-66v Total Bills: %v |\n", "", totalBills)
	fs += "-------------------------------------------------------------------------------------\n"

	fmt.Println(fs)
}

func FindBill(id int64) *Bill {
	b := billsList[id]

	return b
}
