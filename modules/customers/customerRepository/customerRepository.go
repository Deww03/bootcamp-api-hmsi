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

func (db *DB) Create(c *models.RequestInsertCustomer) error {
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

func (db *DB) Update(c *models.RequestUpdateCustomer) error {
	stmt, err := db.Conn.Prepare("UPDATE customers SET name = $1, phone = $2, email = $3, age = $4 WHERE id = $5")
	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)

	if err != nil {
		return errExec
	}

	return nil
}

func (db *DB) Delete(Id uint64) error {
	_, errExec := db.Conn.Exec(`DELETE FROM customers WHERE id = $1`, Id)
	if errExec != nil {
		return errExec
	}

	return nil
}
