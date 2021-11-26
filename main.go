package main

import (
	"fmt"
	"test_task/configs"
	"test_task/model"
)

func main() {
	configs.LoadConfiguration("config.json")
	// fmt.Println(configs.Config)

	// bRes, err := model.BankResponse().GetBankResponse("15.04.2021")
	// fmt.Println(err)
	// fmt.Println(bRes)

	// err = model.Currency().AddRecordToDB(bRes.CurrencyList, bRes.Date)
	// fmt.Println(err)

	res, err := model.Currency().GetByTitleDate("GEL", "15.04.2021")

	fmt.Println(err)
	fmt.Println(res)
}
