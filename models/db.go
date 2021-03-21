package models

var DB []Item

type Item struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}

func init() {
	book1 := Item{
		ID:     1,
		Title:  "New Item",
		Amount: 100,
		Price:  12.55,
	}
	DB = append(DB, book1)
}

func FindItemById(id int) (*Item, bool) {
	var itemIndex int
	var found bool

	for index, b := range DB {
		if b.ID == id {
			itemIndex = index
			found = true
			break
		}
	}

	return &DB[itemIndex], found
}

func RemoveItemById(id int) (*Item, bool) {
	var itemIndex int
	var found bool

	for index, b := range DB {
		if b.ID == id {
			itemIndex = index
			found = true
			break
		}
	}

	var item Item

	if found {
		item = DB[itemIndex]

		var tmpDB = make([]Item, 0, len(DB)-1)
		tmpDB = append(tmpDB, DB[:itemIndex]...)
		tmpDB = append(tmpDB, DB[itemIndex+1:]...)

		DB = tmpDB
	}

	return &item, found
}
