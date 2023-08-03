package customerHandler

import (
	"net/http"
	"strconv"

	"github.com/Deww03/bootcamp-api-hmsi/models"
	"github.com/Deww03/bootcamp-api-hmsi/modules/customers"
	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	UC customers.CustomerUsecase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomerUsecase) {
	handler := customerHandler{UC}

	r.GET("/customers", handler.GetAll)
	r.POST("/customer/create", handler.Insert)
	r.PUT("/customer/update", handler.Update)
	r.DELETE("/customer/delete", handler.Delete)
}

// Implement GetAll
func (h *customerHandler) GetAll(c *gin.Context) {
	result, err := h.UC.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    result,
	})
}

// Implement Insert
func (h *customerHandler) Insert(c *gin.Context) {
	var request models.RequestInsertCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	err := h.UC.Insert(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Inserted successfully",
		"data":    []string{},
	})
}

// Implement Update
func (h *customerHandler) Update(c *gin.Context) {
	var request models.RequestUpdateCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	err := h.UC.Update(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated successfully",
		"data":    []string{},
	})
}

// Implement Delete
func (h *customerHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID cannot be empty",
			"data":    []string{},
		})
		return
	}

	// Convert string to uint64
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID format",
			"data":    []string{},
		})
		return
	}

	// Call Delete usecase.
	err = h.UC.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	// When not error then...
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted successfully",
		"data":    []string{},
	})
}
