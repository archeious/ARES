// Package series provides the interface and repository to create and store
// TV/Anime Series
package series

import (
	"models/item"
)

type Series interface {
	item.Item
	JName() string
	SetJName(string)
	ExtIDs() map[string]string
	SetExtID(string, string)
	Synopsis() string
	SetSynopsis(string)
	//Seasons  []*Season
}

type SeriesRepository interface {
	GetAllSeries() ([]Series, error)
	GetSeriesByName(string) (Series, error)
	GetSeriesById(string) (Series, error)
	NewSeries(string) (Series, error)
	SaveSeries(Series) error
}
