package requests

type ListSkusRequest struct {

	CoffeeMachineProductTypeFilter       	string  `form:"coffee_machine_product_type_filter"`
	CoffeePodProductTypeFilter      		string  `form:"coffee_pod_product_type_filter"`
	WaterLineCompatible        				string  `form:"water_line_compatible_filter"`
	CoffeeFlavorFilter     					string  `form:"coffee_flavor_filter"`
	PackSizeFilter 							string	`form:"pack_size_filter"`
	Model									string	`form:"model"`
}