package series

import (
	"models/item"
)

type ConcreteSeries struct {
	item.BaseItem
	//	Seasons []*Season
}

//TODO: Add error checking
func NewConceteSeries(name, species, id string) (ConcreteSeries, error) {
	return ConcreteSeries{item.NewBaseItem(name, species, id)}, nil
}

/* func (s *ConcreteSeries) String() string {
	return s.name
}

func (s *ConcreteSeries) Name() string {
	return s.name
}

func (s *ConcreteSeries) Id() string {
	return s.id
}
*/
