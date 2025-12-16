package document

import "errors"

type DueAmount struct {
	value int
}

func (a DueAmount) Value() int {
	return a.value
}

func NewDueAmount(value int) (*DueAmount, error) {
	if value < 0 {
		return nil, errors.New("due amount cannot be negative")
	}

	return &DueAmount{value: value}, nil
}

func RehydrateDueAmount(value int) DueAmount {
	return DueAmount{value: value}
}

func (a DueAmount) Subtract(amount int) (DueAmount, error) {
	if amount < 0 {
		return DueAmount{}, errors.New("amount to subtract cannot be negative")
	}

	if amount > a.value {
		return DueAmount{}, errors.New("amount to subtract cannot be greater than due amount")
	}

	return DueAmount{value: a.value - amount}, nil
}
