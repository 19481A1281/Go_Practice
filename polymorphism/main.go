package main

import "fmt"

type purchasable interface{
	CalculatePrice() int64
}

var cart []purchasable

func addToCart(products ...purchasable){
	cart = append(cart, products...)
}

func getCartTotal() int64{
	var total int64

	for _, product := range cart{
		total+=product.CalculatePrice()
		fmt.Println(total)
	}

	return total
}

func main(){
	shirt := Shirt{Product{100, "NIKE"}, "M", "YELLOW"}
	monitor := Monitor{Product{200, "DELL"},"High", 55}

	addToCart(shirt, monitor)
	fmt.Println(getCartTotal())

}