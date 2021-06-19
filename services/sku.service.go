package services

import (
	"github.com/swnsonhe/task/repositories"
	"github.com/swnsonhe/task/resources"
	"github.com/swnsonhe/task/requests" 

)
 
type SKUService struct {
	repo *repositories.SKURepository
}

func NewSKUService() *SKUService {
	service := SKUService{
		repo: repositories.GetSKURepository(),
	}
	return &service
}
 
func (service SKUService) Find(listSkusRequest requests.ListSkusRequest)  resources.IResource {

	SKUs, err  := service.repo.Find(listSkusRequest)

	if err != nil {
		resource := resources.GetError500Resource("something went wrong!")
		return resource
	}
	
	return resources.GetSuccess200Resource(SKUs, "")
}
 
