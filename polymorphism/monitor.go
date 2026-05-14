package main

type Monitor struct{
	Product
	Resolution string
	Width int64
}

func (m Monitor) CalculatePrice() int64{
	tax := float64(m.Price) * 0.30
	return m.Price + int64(tax)
}