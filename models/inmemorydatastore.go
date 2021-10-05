package models

type InmemoryDatastore struct{
	Products map[int]Product
}

func InitialiseInmemoryDatastore() *InmemoryDatastore {
	products:=make(map[int]Product)
	products[1]= Product{Id: 1, Name: "bat", Description: "mrf bat", Price: 220, Quantity: 10}
	products[2]= Product{Id: 2, Name: "stump", Description: "mrf stump", Price: 22, Quantity: 10}
	products[3]= Product{Id: 3, Name: "helmet", Description: "solid helmet", Price: 220, Quantity: 10}
	p:= InmemoryDatastore{products}
	return &p
}
