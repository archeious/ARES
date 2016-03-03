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
	return &ConcreteSeries{item.NewBaseItem(name, "", id)}, err
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
		itemLst = append(itemLst, &ConcreteSeries{item.NewBaseItem(name, "", id)})
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
	return &ConcreteSeries{item.NewBaseItem(name, "", id)}, err
}

func (i *MysqlSeriesRepository) NewSeries(name string, species string) (Series, error) {
	u, err := uuid.NewV4()
	newSeries := ConcreteSeries{item.NewBaseItem(name, species, u.String())}

	stmt, err := i.db.Prepare("INSERT INTO series(id,name) VALUES (?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(u.String(), name)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			switch {
			case driverErr.Number == 1062: // Item already existrs: http://dev.mysql.com/doc/refman/5.5/en/error-messages-server.html#error_er_dup_entry
				return nil, item.ErrAlreadyExistsInRepo
			default:
				return nil, err
			}
		}
	}
	return &newSeries, nil
}

//TODO: Add error handling
func NewMysqlSeriesRepository(dbh *sql.DB) (*MysqlSeriesRepository, error) {
	return &MysqlSeriesRepository{db: dbh}, nil
}

func init() {
	idCount = 1
	series = make(map[string]ConcreteSeries)
}
