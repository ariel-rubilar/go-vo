package document

import (
	"errors"
	"slices"
)

const (
	statusPendingValue = "pending"
	statusPaidValue    = "paid"
	statusOverdueValue = "overdue"
)

var (
	validStatusValues = []string{statusPendingValue, statusPaidValue, statusOverdueValue}
	StatusPending     = Status{value: statusPendingValue}
	StatusPaid        = Status{value: statusPaidValue}
	StatusOverdue     = Status{value: statusOverdueValue}
)

type Status struct {
	value string
}

func RehydrateStatus(value string) Status {
	return Status{value: value}
}

func NewStatus(value string) (*Status, error) {

	if !slices.Contains(validStatusValues, value) {
		return nil, errors.New("invalid status value")
	}

	return &Status{value: value}, nil
}

func (s Status) Value() string {
	return s.value
}

func (s Status) IsUnknown() bool {
	return s != StatusPending && s != StatusPaid && s != StatusOverdue
}

func (s Status) Is(status Status) bool {
	return s == status
}

func (s Status) SomeOf(statuses ...Status) bool {
	return slices.Contains(statuses, s)
}

func (s Status) IsAlive() bool {
	return s == StatusPending || s == StatusOverdue
}

func (s Status) IsFinal() bool {
	return s == StatusPaid
}

func (s Status) IsPending() bool {
	return s == StatusPending
}

func (s Status) IsPaid() bool {
	return s == StatusPaid
}

func (s Status) IsOverdue() bool {
	return s == StatusOverdue
}
