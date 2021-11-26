package model

import (
	"test_task/configs"
)

type CurrencyModel struct {
	Id    int     `db:"id"`
	Title string  `xml:"fullname" db:"title" json:"currency"`
	Code  string  `xml:"title" db:"code" json:"code"`
	Value float64 `xml:"description" db:"value" json:"value"`
	Date  string  `db:"a_date"`
}

func Currency() *CurrencyModel {
	return &CurrencyModel{}
}

func (*CurrencyModel) AddRecordToDB(currencies []CurrencyModel, date string) error {
	query := `insert into r_currency (title, code, value, a_date)
	VALUES ($1, $2, $3, to_date($4, 'DD.MM.YYYY'))
	 on conflict (code) do update set
	 title = $1, code = $2, value = $3, a_date = to_date($4, 'DD.MM.YYYY')`

	for _, currency := range currencies {
		_, err := configs.DB().Exec(query, currency.Title, currency.Code, currency.Value, date)

		if err != nil {
			return err
		}
	}
	return nil
}

func (*CurrencyModel) GetByTitleDate(code, date string) (CurrencyModel, error) {
	res := CurrencyModel{}

	query := `select * from r_currency where code = $1 and a_date = to_date($2, 'DD.MM.YYYY')`

	err := configs.DB().Get(&res, query, code, date)

	return res, err
}
