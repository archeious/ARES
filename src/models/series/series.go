package series

import (
	"models/item"
)

type Series interface {
	item.Item
	//Seasons  []*Season
}

type SeriesRepository interface {
	GetSeriesByName(string) (Series, error)
	GetSeriesById(string) (Series, error)
	NewSeries(string, string) (Series, error)
}
