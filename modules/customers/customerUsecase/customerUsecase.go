package customerUsecase

import (
	"github.com/Deww03/bootcamp-api-hmsi/models"
	"github.com/Deww03/bootcamp-api-hmsi/modules/customers"
)

type customerRepository struct {
	Repo customers.CustomerRepository
}

func NewCustomerUsecase(Repo customers.CustomerRepository) customers.CustomerUsecase {
	return &customerRepository{Repo}
}

func (r *customerRepository) FindAll() (*[]models.Customers, error) {
	result, err := r.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (r *customerRepository) Insert(c *models.Customers)error {
// 	err := r.Repo.Create(c)

// 	if err!= nil {
//         return err
//     }
// 	return nil
// }
