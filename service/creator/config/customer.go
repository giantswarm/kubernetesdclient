package config

func NewCustomer() *Customer {
	return &Customer{
		Name: "",
	}
}

type Customer struct {
	Name string
}
