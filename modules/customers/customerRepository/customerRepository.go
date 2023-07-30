package customerRepository

import (
	"database/sql"
	"fmt"

	"github.com/Deww03/bootcamp-api-hmsi/models"
	"github.com/Deww03/bootcamp-api-hmsi/modules/customers"
)

type DB struct {
	Conn *sql.DB
}

func NewCustomerRepository(Conn *sql.DB) customers.CustomerRepository {
	return &DB{Conn}
}

func (db *DB) GetAll() (*[]models.Customers, error) {
	rows, errExac := db.Conn.Query(`select id, name, phone, email, age FROM customers`)
	if errExac != nil {
		return nil, errExac
	}

	// deklarasi variabel result
	var result []models.Customers

	for rows.Next() {
		var each models.Customers

		errScan := rows.Scan(&each.Id, &each.Name, &each.Phone, &each.Email, &each.Age)

		if errScan != nil {
			return nil, errScan
		}

		result = append(result, each)
	}

	return &result, nil
}

func (db *DB) Create(c *models.Customers) error {
	stmt, err := db.Conn.Prepare("INSERT INTO customers (name, phone, email, age) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)

	if err != nil {
		fmt.Println("errExec", errExec)
		return errExec
	}

	return nil
}
