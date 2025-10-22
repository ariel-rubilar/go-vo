package document

import (
	"errors"
	"slices"
)

type DueAmount struct {
	value int
}

func (a DueAmount) Value() int {
	return a.value
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

func NewDueAmount(value, total int) (*DueAmount, error) {
	if value < 0 {
		return nil, errors.New("due amount cannot be negative")
	}

	if value > total {
		return nil, errors.New("due amount cannot be greater than total amount")
	}

	return &DueAmount{value: value}, nil
}

func RehydrateDueAmount(value int) DueAmount {
	return DueAmount{value: value}
}

var (
	StatusPending = Status{value: "pending"}
	StatusPaid    = Status{value: "paid"}
	StatusOverdue = Status{value: "overdue"}
)

type Status struct {
	value string
}

func RehydrateStatus(value string) Status {
	return Status{value: value}
}

func NewStatus(value string) (*Status, error) {
	if value != StatusPending.Value() && value != StatusPaid.value && value != StatusOverdue.value {
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

type Document struct {
	Amount    int
	DueAmount DueAmount
	Status    Status
}

func (d *Document) AddPayment(payment int) error {

	newDue, err := d.DueAmount.Subtract(payment)

	if err != nil {
		return err
	}

	d.DueAmount = newDue

	if newDue.Value() == 0 {
		d.Status = StatusPaid
	} else if newDue.Value() < d.Amount {
		d.Status = StatusPending
	}

	return nil
}

func New(amount int, dueAmount DueAmount, status Status) (*Document, error) {
	if amount < 0 {
		return nil, errors.New("amount cannot be negative")
	}

	if status != StatusPending && status != StatusPaid && status != StatusOverdue {
		return nil, errors.New("invalid status")
	}

	return &Document{
		Amount:    amount,
		DueAmount: dueAmount,
		Status:    status,
	}, nil
}

func InitDocument(amount int) (*Document, error) {

	dueAmount, err := NewDueAmount(amount, amount)

	if err != nil {
		return nil, err
	}
	return &Document{
		Amount:    amount,
		DueAmount: *dueAmount,
		Status:    StatusPending,
	}, nil
}

func Rehydrate(amount, dueAmount int, status string) *Document {
	dueAmt := RehydrateDueAmount(dueAmount)
	s := RehydrateStatus(status)

	return &Document{
		Amount:    amount,
		DueAmount: dueAmt,
		Status:    s,
	}
}

func NewFromPrimitives(amount int, dueAmount int, status string) (*Document, error) {
	dueAmt, err := NewDueAmount(dueAmount, amount)
	if err != nil {
		return nil, err
	}

	s, err := NewStatus(status)
	if err != nil {
		return nil, err
	}

	return New(amount, *dueAmt, *s)
}

type ResponseID struct {
	value *int
}

func NewResponseID(value *int) ResponseID {

	if value == nil {
		return ResponseID{value: nil}
	}

	return ResponseID{value: value}
}

func (r ResponseID) Value() *int {
	return r.value
}

func (r ResponseID) IsEmpty() bool {
	return r.value == nil || *r.value <= 0
}

type Response struct {
	ID      ResponseID
	Message string
}

func (r *Response) IsSuccess() bool {
	return !r.ID.IsEmpty()
}

func (r *Response) IsFailure() bool {
	return r.ID.IsEmpty()
}
