package bill

import (
	"time"

	"github.com/onebitgod/billsystem/item"
)

type Bill struct {
	ID           int64
	Items        []BillItem
	CustomerName string
	Tip          float64
	Total        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type BillItem struct {
	Item  item.Item
	Count int
}
