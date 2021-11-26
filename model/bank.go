package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"test_task/configs"
)

type BankResponseModel struct {
	Date         string          `xml:"date"`
	CurrencyList []CurrencyModel `xml:"item"`
}

func BankResponse() *BankResponseModel {
	return &BankResponseModel{}
}

func (*BankResponseModel) GetBankResponse(date string) (BankResponseModel, error) {
	res := BankResponseModel{}

	// url + date
	url := configs.Config.Url + date
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return res, err
	}

	fmt.Println("client")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return res, err
	}

	err = xml.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	return res, err
}
