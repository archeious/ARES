package item

type BaseItem struct {
	id   string
	name string
	// why species and not type?  Well if you have syntax highlighting on you will see why.
	// I can't use "type" and I am in a weird mood, so suck it future self.
	species string
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
