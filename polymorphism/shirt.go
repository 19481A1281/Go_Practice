package main

type Shirt struct{
	Product
	Size string
	Color string
}

func (s Shirt) CalculatePrice() int64{
	discount := float64(s.Price) * 0.20
	return s.Price-int64(discount)
}