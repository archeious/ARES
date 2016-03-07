package item

type BaseItem struct {
	id   string `schema:"id"`
	name string `schema:"name"`
	// why species and not type?  Well if you have syntax highlighting on you will see why.
	// I can't use "type" and I am in a weird mood, so suck it future self.
	species string `'schema:"species"`
}

func (i *BaseItem) Name() string {
	return i.name
}

func (i *BaseItem) Id() string {
	return i.id
}

func (i *BaseItem) Species() string {
	return i.species
}

func (i *BaseItem) SetName(newName string) {
	i.name = newName
}

func NewBaseItem(name, species, id string) BaseItem {
	return BaseItem{name: name, species: species, id: id}
}
