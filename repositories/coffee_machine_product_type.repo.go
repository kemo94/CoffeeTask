package repositories

import (
	"github.com/swnsonhe/task/models"  
	"github.com/swnsonhe/task/requests"  

)
 
type SKURepository struct {
	productType models.CoffeeMachineProductTypeModel
}
 
func GetSKURepository() *SKURepository {
	return &SKURepository{}
}

func (repo SKURepository) Find(listSkusRequest requests.ListSkusRequest) (skus []models.SKU, err error) {
	skus, err = models.GetSKUS(listSkusRequest)

	return skus, err
}
