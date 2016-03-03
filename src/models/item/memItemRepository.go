package item

import "errors"

var (
	idCount int64 = 0
	items   map[string]BaseItem
)

type MemItemRepository struct {
}

func (i *MemItemRepository) GetAll() []Item {
	itemLst := make([]Item, 0)
	for _, item := range items {
		itemLst = append(itemLst, &item)
	}
	return itemLst
}

func (i *MemItemRepository) GetByName(n string) (Item, error) {
	return &BaseItem{}, errors.New("memItemRepository: GetItemByName not implemented!")
}

func (i *MemItemRepository) GetById(id string) (Item, error) {
	if i, ok := items[id]; !ok {
		return &BaseItem{}, errors.New("memItemRepository: Item ID " + id + "does not exist!")
	} else {
		return &i, nil
	}
}

func (i *MemItemRepository) NewItem(name string, species string) (Item, error) {
	newItem := BaseItem{name: name, species: species, id: string(idCount)}
	items[string(idCount)] = newItem
	idCount += 1
	return &newItem, nil
}

func init() {
	idCount = 1
	items = make(map[string]BaseItem)
}
