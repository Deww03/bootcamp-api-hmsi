package customers

import "github.com/Deww03/bootcamp-api-hmsi/models"

type (
	CustomerRepository interface {
		GetAll() (*[]models.Customers, error)
		Create(c *models.RequestInsertCustomer) error
		Update(c *models.RequestUpdateCustomer) error
		Delete(id uint64) error
	}

	CustomerUsecase interface {
		FindAll() (*[]models.Customers, error)
		Insert(c *models.RequestInsertCustomer) error
		Update(c *models.RequestUpdateCustomer) error
		Delete(id uint64) error
	}
)
