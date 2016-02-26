package item

type BaseItem struct {
	id   string
	name string
}

func (i *BaseItem) Name() string {
	return i.name
}

func (i *BaseItem) Id() string {
	return i.id
}

func (i *BaseItem) SetName(newName string) {
	i.name = newName
}
