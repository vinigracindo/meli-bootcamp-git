package main

import (
	"github.com/vinigracindo/meli-bootcamp-git/model"
)

func main() {
	company := model.NewCompany("Mercado Livre 2")
	//company.NewSales("1,14.58,5;2,8.99,2;")
	company.TotalSales()
}
