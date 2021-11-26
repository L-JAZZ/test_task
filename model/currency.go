package model

import (
	"time"
)

type CurrencyModel struct {
	Title string  `xml:"fullname" db:"title" json:"currency"`
	Code  string  `xml:"title" db:"code" json:"code"`
	Value float64 `xml:"description" db:"value" json:"value"`
}

func Currency() *CurrencyModel {
	return &CurrencyModel{}
}

func (*CurrencyModel) AddRecordToDB(currencies []CurrencyModel) error {
	query := `insert into r_currency (title, code, value, a_date)
	VALUES ($1, $2, $3, $4)
	 on conflict (code) do update set
	 title = $1, code = $2, value = $3, a_date = $4`

	currentTime := time.Now()

	for _, currency := range currencies {
		_, err := someshit
	}
}
