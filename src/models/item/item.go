package item

type Item interface {
	Name() string
	SetName(string)
	Id() string
}

type ItemRepository interface {
	GetItemByName(string) (Item, error)
	GetItemById(string) (Item, error)
	NewItem(string) (Item, error)
}
