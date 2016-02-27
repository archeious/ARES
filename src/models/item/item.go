package item

type Item interface {
	Name() string
	SetName(string)
	Id() string
	Species() string
}

type ItemRepository interface {
	GetItemByName(string) (Item, error)
	GetItemById(string) (Item, error)
	NewItem(string, string) (Item, error)
}
