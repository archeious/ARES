package series

import (
	"models/item"
)

type ConcreteSeries struct {
	item.BaseItem
	jName  string
	extIDs map[string]string
	//	Seasons []*Season
}

//TODO: Add error checking
func NewConceteSeries(name, species, id string) (ConcreteSeries, error) {
	return ConcreteSeries{item.NewBaseItem(name, species, id), "", nil}, nil
}

func (s *ConcreteSeries) JName() string {
	return s.jName
}

func (s *ConcreteSeries) ExtIDs() map[string]string {
	return s.extIDs
}

func (s *ConcreteSeries) SetExtID(key, id string) {
	if s.extIDs == nil {
		s.extIDs = make(map[string]string)
	}
	s.extIDs[key] = id
}
