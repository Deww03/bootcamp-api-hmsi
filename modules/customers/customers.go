package customers

import "github.com/Deww03/bootcamp-api-hmsi/models"

type (
	CustomerRepository interface {
		GetAll() (*[]models.Customers, error)
	}

	CustomerUsecase interface {
		FindAll() (*[]models.Customers, error)
	}
)
