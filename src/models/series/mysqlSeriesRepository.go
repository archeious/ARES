package series

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/nu7hatch/gouuid"
	"log"
	"models/item"
)

type MysqlSeriesRepository struct {
	db *sql.DB
}

func (i *MysqlSeriesRepository) GetSeriesByName(n string) (Series, error) {
	var id string
	var name string
	query := "select id, name from series where name = ?"

	err := i.db.QueryRow(query, n).Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, item.ErrDoesNotExist
		} else {
			log.Fatal(err)
		}
	}
	return &ConcreteSeries{BaseItem: item.NewBaseItem(name, "", id)}, err
}

func (i *MysqlSeriesRepository) GetAllSeries() ([]Series, error) {
	itemLst := make([]Series, 0)
	query := "select id, name from series"
	rows, err := i.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		itemLst = append(itemLst, &ConcreteSeries{BaseItem: item.NewBaseItem(name, "", id)})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return itemLst, nil
}

func (i *MysqlSeriesRepository) GetAll() ([]item.Item, error) {
	itemLst := make([]item.Item, 0)
	seriesLst, _ := i.GetAllSeries()
	for _, v := range seriesLst {
		itemLst = append(itemLst, v)
	}
	//TODO: error check
	return itemLst, nil
}

func (i *MysqlSeriesRepository) GetSeriesById(id string) (Series, error) {
	var name string
	query := "select id, name from series where id = ?"
	log.Println(query, id)
	err := i.db.QueryRow(query, id).Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, item.ErrDoesNotExist
		} else {
			log.Fatal(err)
		}
	}

	s := ConcreteSeries{BaseItem: item.NewBaseItem(name, "", id)}
	i.GetSeriesExtIds(&s)
	return &s, err
}

func (i *MysqlSeriesRepository) NewSeries(name string, species string) (Series, error) {
	u, err := uuid.NewV4()
	newSeries := ConcreteSeries{BaseItem: item.NewBaseItem(name, species, u.String())}

	stmt, err := i.db.Prepare("INSERT INTO series(id,name) VALUES (?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(u.String(), name)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			switch {
			case driverErr.Number == 1062: // Item already existrs: http://dev.mysql.com/doc/refman/5.5/en/error-messages-server.html#error_er_dup_entry
				//TODO: Error check
				s, _ := i.GetSeriesByName(name)
				return s, item.ErrAlreadyExists
			default:
				return nil, err
			}
		}
	}
	return &newSeries, nil
}

func (i *MysqlSeriesRepository) SaveSeries(s Series) error {
	log.Println("saving:", s)
	tx, err := i.db.Begin()
	log.Println("deleting ids:", s)
	stmt, err := i.db.Prepare("delete from extId where id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err = stmt.Exec(s.Id()); err != nil {
		tx.Rollback()
		return err
	}

	if extIdStmt, err := i.db.Prepare("insert into extId (id,name,extId) values (?,?,?)"); err != nil {
		tx.Rollback()
		return err
	} else {
		log.Println("inserting ids:", s.ExtIDs())
		for k, v := range s.ExtIDs() {
			log.Println("inserting s.Id(),k,v:", s.Id(), k, v)
			if _, err = extIdStmt.Exec(s.Id(), k, v); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func (i *MysqlSeriesRepository) GetSeriesExtIds(s Series) {
	query := "select name, extId from extId where id = ?"
	rows, err := i.db.Query(query, s.Id())
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var extid string
		var name string
		err := rows.Scan(&name, &extid)
		if err != nil {
			log.Println(err)
			return
		}
		s.SetExtID(name, extid)
	}
	err = rows.Err()
	if err != nil {
		return
	}

}

//TODO: Add error handling
func NewMysqlSeriesRepository(dbh *sql.DB) (*MysqlSeriesRepository, error) {
	return &MysqlSeriesRepository{db: dbh}, nil
}

func init() {
	idCount = 1
	series = make(map[string]ConcreteSeries)
}
