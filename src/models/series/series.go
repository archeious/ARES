package series

import (
	"log"
	"models/item"
)

type Series interface {
	Item
	//Seasons  []*Season
}

type SeriesRepository interface {
	ItemRepository
}
