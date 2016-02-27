package item

import "errors"

var (
	idCount int64 = 0
	items   map[string]BaseItem
)

type memItemRepository struct {
}

func (i *memItemRepository) GetItemByName(n string) (Item, error) {
	return &BaseItem{}, errors.New("memItemRepository: GetItemByName not implemented!")
}

func (i *memItemRepository) GetItemById(id string) (Item, error) {
	if i, ok := items[id]; !ok {
		return &BaseItem{}, errors.New("memItemRepository: Item ID " + id + "does not exist!")
	} else {
		return &i, nil
	}
}

func (i *memItemRepository) NewItem(name string, species string) (Item, error) {
	newItem := BaseItem{name: name, species: species, id: string(idCount)}
	items[string(idCount)] = newItem
	idCount += 1
	return &newItem, nil
}

func init() {
	idCount = 1
	items = make(map[string]BaseItem)
}
