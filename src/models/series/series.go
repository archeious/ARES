// Package series provides the interface and repository to create and store
// TV/Anime Series
package series

import (
	"models/item"
)

type Series interface {
	item.Item
	//Seasons  []*Season
}

type SeriesRepository interface {
	GetAllSeries() ([]Series, error)
	GetSeriesByName(string) (Series, error)
	GetSeriesById(string) (Series, error)
	NewSeries(string, string) (Series, error)
}
