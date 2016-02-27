package models

import (
	"github.com/revel/revel"
	"log"
	"webfrontend/app"
)

type Season struct {
	Id       int64
	Name     string
	Series   *Series
	Episodes *[]Episode
	Number   string
}

func (s *Season) String() string {
	return s.Name
}

func (s *Season) Validate(v *revel.Validation) {
	v.Check(s.Name,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{255},
	)
}

func (s *Season) Add() {
	const query = "INSERT INTO season (name, series_id) VALUES (?,?)"

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(s.Name, s.Series.Id)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	s.Id = int64(id)
}

func GetSeasonsBySeriesId(seriesId int64) []Season {

	const query = "SELECT id, name from season where series_id = ?"

	var name string
	var id int64
	seasons := make([]Season, 0)

	rows, err := app.DB.Query(query, seriesId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		seasons = append(seasons, Season{Name: name, Id: id})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return seasons
}
