package series

import (
	"errors"
	"models/item"
)

var (
	idCount int64 = 0
	series  map[string]ConcreteSeries
)

type MemSeriesRepository struct {
}

func (i *MemSeriesRepository) GetSeriesByName(n string) (Series, error) {
	return nil, errors.New("memSeriesRepository: GetSeriesByName not implemented!")
}

func (i *MemSeriesRepository) GetSeriesById(id string) (Series, error) {
	if i, ok := series[id]; !ok {
		return nil, errors.New("memSeriesRepository: Series ID " + id + "does not exist!")
	} else {
		return &i, nil
	}
}

func (i *MemSeriesRepository) NewSeries(name string, species string) (Series, error) {
	newSeries := ConcreteSeries{item.NewBaseItem(name, species, string(idCount))}
	series[string(idCount)] = newSeries
	idCount += 1
	return &newSeries, nil
}

func init() {
	idCount = 1
	series = make(map[string]ConcreteSeries)
}
