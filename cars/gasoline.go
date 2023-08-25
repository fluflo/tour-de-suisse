package cars

type Gasoline struct {
	name  string
	price float64 // price per liters in chf
}

func NewGasoline(name string, price float64) Gasoline {
	return Gasoline{name, price}
}
