package connectDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnPostgres(host, port, user, password, dbname, driver string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open(driver, psqlInfo)

	if err != nil {
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
