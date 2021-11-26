package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"test_task/model"

	"github.com/gorilla/mux"
)

func SaveAPI(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	bankRes, err := model.BankResponse().GetBankResponse(date)

	if err != nil {
		log.Println(err)

		jsonResponse, jsonError := json.Marshal(CommonResponse{
			Success: false,
			Data:    nil,
			Message: "Failed to get bank response",
		})

		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
	}

	err = model.Currency().AddRecordToDB(bankRes.CurrencyList, bankRes.Date)
	if err != nil {
		log.Println(err)

		jsonResponse, jsonError := json.Marshal(CommonResponse{
			Success: false,
			Data:    nil,
			Message: "Failed to save currencies",
		})

		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
	}

	jsonResponse, jsonError := json.Marshal(CommonResponse{
		Success: true,
		Data:    bankRes,
		Message: "Success",
	})
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetCurrency(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	code := params["code"]

	res, err := model.Currency().GetByTitleDate(code, date)
	if err != nil {
		log.Println(err)

		jsonResponse, jsonError := json.Marshal(CommonResponse{
			Success: false,
			Data:    nil,
			Message: "Failed to get currency",
		})

		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
	}

	jsonResponse, jsonError := json.Marshal(CommonResponse{
		Success: false,
		Data:    res,
		Message: "Success",
	})

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
