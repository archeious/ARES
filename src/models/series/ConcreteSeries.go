package series

type ConcreteSeries struct {
	name string
	id   string
	//	Seasons []*Season
}

func (s *Series) String() string {
	return s.name
}

func (s *Series) Name() string {
	return s.name
}

func (S *Series) Id() string {
	return id
}
