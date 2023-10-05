package types

import (
	"github.com/shopspring/decimal"
)

type Currency struct {
	value decimal.Decimal
}

func NewCurrency(value decimal.Decimal) Currency {
	return Currency{value: value}
}

func CurrencyFromString(str string) (Currency, error) {
	value, err := decimal.NewFromString(str)
	if err != nil {
		return Currency{}, err
	}
	return NewCurrency(value), nil
}

func (c Currency) String() string {
	return c.value.String()
}

func (c Currency) Decimal() decimal.Decimal {
	return c.value
}
