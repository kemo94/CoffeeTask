package models

import (
	"fmt"
	"github.com/swnsonhe/task/requests"  
	"strconv"

)
type SKU struct {
	Id        								string  `json:"id"`
	CoffeeMachineProductTypeModelId        	string  `json:"coffeeMachineProductTypeModelId"`
	CoffeePodProductTypeModelId        		string  `json:"coffeePodProductTypeModelId"`
	WaterLineCompatible        				bool  	`json:"waterLineCompatible"`
	CoffeeFlavorId        					string  `json:"coffeeFlavorId"`
	Model        							string  `json:"model"`
	PackSizeId 								string	`json:"packSizeId"`
}


func GetSKUS(filter requests.ListSkusRequest ) ([]SKU, error) {
	dummyDataSkus := DummyData()

	skus := []SKU{}
	for _, sku := range dummyDataSkus {
		
		var  checkCoffeeMachineProductType,checkCoffeePodProductType,checkCoffeeFlavor,checkWaterLine,checkPackSize,checkSkuModel bool

		if checkCoffeeMachineProductType = false; sku.CoffeeMachineProductTypeModelId == filter.CoffeeMachineProductTypeFilter || len(filter.CoffeeMachineProductTypeFilter) == 0  {
			checkCoffeeMachineProductType = true 
		} 

		if checkCoffeePodProductType = false; sku.CoffeePodProductTypeModelId == filter.CoffeePodProductTypeFilter || len(filter.CoffeePodProductTypeFilter) == 0  {
			checkCoffeePodProductType = true 
		} 


		if checkCoffeeFlavor = false; sku.CoffeeFlavorId == filter.CoffeeFlavorFilter || len(filter.CoffeeFlavorFilter) == 0  {
			checkCoffeeFlavor = true 
		} 

		check, err := strconv.ParseBool(filter.WaterLineCompatible)
		if err != nil {
			checkWaterLine = true
		}else{
			if  sku.WaterLineCompatible == check {
				checkWaterLine = true
			}else{
				checkWaterLine = false
			}
		}
		

		if checkPackSize = false; sku.PackSizeId == filter.PackSizeFilter || len(filter.PackSizeFilter) == 0  {
			checkPackSize = true 
		} 


		if checkSkuModel = false; sku.Model == filter.Model || len(filter.Model) == 0  {
			checkSkuModel = true
		}
		if ( sku.Id ==  "CM003"){
			fmt.Println("cpmare" ,checkWaterLine, check,  sku.WaterLineCompatible, err);
		}
		if ( checkCoffeeMachineProductType &&
			 checkCoffeePodProductType &&
			 checkWaterLine &&
			 checkCoffeeFlavor &&
			 checkPackSize &&
			 checkSkuModel){
			
				skus = append(skus, sku )
		} 	
	}


	// CoffeeMachineProductTypeFilter       	string  `form:"coffee_machine_product_type_filter"`
	// CoffeePodProductTypeFilter      		string  `form:"coffeePodProductTypeFilter"`
	// WaterLineCompatible        				bool  	`form:"waterLineCompatibleFilter"`
	// CoffeeFlavorFilter     					string  `form:"coffeeFlavorFilter"`
	// PackSizeFilter 	

	return skus,nil
}
	



func DummyData() ([]SKU) {

	skus := []SKU{}
	skus = append(skus, SKU{
		Id : "CM001" , CoffeeMachineProductTypeModelId : "COFFEE_MACHINE_SMALL", Model : "base",
	})
	skus = append(skus, SKU{
		Id : "CM002" , CoffeeMachineProductTypeModelId : "COFFEE_MACHINE_SMALL", Model : "premium",
	})
	skus = append(skus, SKU{
		Id : "CM003" , CoffeeMachineProductTypeModelId : "COFFEE_MACHINE_SMALL", Model : "deluxe",
		WaterLineCompatible : true,
	})
	skus = append(skus, SKU{
		Id : "CM101" , CoffeeMachineProductTypeModelId : "COFFEE_MACHINE_LARGE", Model : "base",
	})
	skus = append(skus, SKU{
		Id : "CM102" , CoffeeMachineProductTypeModelId : "COFFEE_MACHINE_LARGE", Model : "premium",
		WaterLineCompatible : true,
	})
	skus = append(skus, SKU{
		Id : "CM103" , CoffeeMachineProductTypeModelId : "COFFEE_MACHINE_LARGE", Model : "deluxe",
		WaterLineCompatible : true,
	})
	skus = append(skus, SKU{
		Id : "EM001" , CoffeeMachineProductTypeModelId : "ESPRESSO_MACHINE", Model : "base",
	})
	skus = append(skus, SKU{
		Id : "EM002" , CoffeeMachineProductTypeModelId : "ESPRESSO_MACHINE", Model : "premium",
	})
	skus = append(skus, SKU{
		Id : "EM003" , CoffeeMachineProductTypeModelId : "ESPRESSO_MACHINE", Model : "deluxe",
	})
	skus = append(skus, SKU{
		Id : "CP001" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_VANILLA",
	})
	skus = append(skus, SKU{
		Id : "CP003" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_VANILLA",
	})
	skus = append(skus, SKU{
		Id : "CP011" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_CARAMEL",
	})
	skus = append(skus, SKU{
		Id : "CP013" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_CARAMEL",
	})
	skus = append(skus, SKU{
		Id : "CP021" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_PSL",
	})
	skus = append(skus, SKU{
		Id : "CP023" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_PSL",
	})
	skus = append(skus, SKU{
		Id : "CP031" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_MOCHA",
	})
	skus = append(skus, SKU{
		Id : "CP033" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_MOCHA",
	})
	skus = append(skus, SKU{
		Id : "CP041" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_HAZELNUT",
	})
	skus = append(skus, SKU{
		Id : "CP043" , CoffeePodProductTypeModelId : "COFFEE_POD_SMALL",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_HAZELNUT",
	})

	skus = append(skus, SKU{
		Id : "CP101" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_VANILLA",
	})
	skus = append(skus, SKU{
		Id : "CP103" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_VANILLA",
	})
	skus = append(skus, SKU{
		Id : "CP111" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_CARAMEL",
	})
	skus = append(skus, SKU{
		Id : "CP113" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_CARAMEL",
	})
	skus = append(skus, SKU{
		Id : "CP121" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_PSL",
	})
	skus = append(skus, SKU{
		Id : "CP123" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_PSL",
	})
	skus = append(skus, SKU{
		Id : "CP131" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_MOCHA",
	})
	skus = append(skus, SKU{
		Id : "CP133" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_MOCHA",
	})
	skus = append(skus, SKU{
		Id : "CP141" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "1_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_HAZELNUT",
	})
	skus = append(skus, SKU{
		Id : "CP143" , CoffeePodProductTypeModelId : "COFFEE_POD_LARGE",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_HAZELNUT",
	})


	
	skus = append(skus, SKU{
		Id : "EP003" , CoffeePodProductTypeModelId : "ESPRESSO_POD",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_VANILLA",
	})
	skus = append(skus, SKU{
		Id : "EP005" , CoffeePodProductTypeModelId : "ESPRESSO_POD",
		PackSizeId : "5_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_VANILLA",
	})
	skus = append(skus, SKU{
		Id : "EP007" , CoffeePodProductTypeModelId : "ESPRESSO_POD",
		PackSizeId : "7_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_VANILLA",
	})
	skus = append(skus, SKU{
		Id : "EP013" , CoffeePodProductTypeModelId : "ESPRESSO_POD",
		PackSizeId : "3_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_CARAMEL",
	})
	skus = append(skus, SKU{
		Id : "EP015" , CoffeePodProductTypeModelId : "ESPRESSO_POD",
		PackSizeId : "5_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_CARAMEL",
	})
	skus = append(skus, SKU{
		Id : "EP017" , CoffeePodProductTypeModelId : "ESPRESSO_POD",
		PackSizeId : "7_DOZEN" , CoffeeFlavorId : "COFFEE_FLAVOR_CARAMEL",
	})
	
	return skus
}



