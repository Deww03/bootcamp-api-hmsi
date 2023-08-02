package models

type (
	Customers struct {
		Id    uint
		Name  string
		Phone string
		Email string
		Age   uint
	}

	RequestInsertCustomer struct {
		Name  string `json:"name" binding:"required"`
		Phone string `json:"phone" binding:"required"`
		Email string `json:"email" binding:"required"`
		Age   uint   `json:"age" binding:"required"`
	}

	RequestUpdateCustomer struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Email string `json:"email"`
		Age   uint   `json:"age"`
	}
)
