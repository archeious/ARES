package item

type Item interface {
	Name() string
	SetName(string)
	Id() string
	Species() string
}

type ItemRepository interface {
	GetByName(string) (Item, error)
	GetById(string) (Item, error)
	NewItem(string, string, string) (Item, error)
}
