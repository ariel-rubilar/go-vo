package document

import (
	"errors"
)

type Document struct {
	Amount    int
	DueAmount DueAmount
	Status    Status
	events    []any
}

type DocumentPrimitives struct {
	Amount    int
	DueAmount int
	Status    string
}

func new(amount int, dueAmount DueAmount, status Status) (*Document, error) {

	if dueAmount.value > amount {
		return nil, errors.New("due amount cannot be greater than amount")
	}

	d := &Document{
		Amount:    amount,
		DueAmount: dueAmount,
		Status:    status,
		events:    []any{},
	}

	return d, nil
}

func Create(amount int) (*Document, error) {

	dueAmount, err := NewDueAmount(amount)

	if err != nil {
		return nil, err
	}

	d, err := new(amount, *dueAmount, StatusPending)

	if err != nil {
		return nil, err
	}

	primitives := d.ToPrimitives()

	// TODO: Implement event recording
	d.recordEvent(struct {
		Amount    int
		DueAmount int
		Status    string
	}{
		Amount:    primitives.Amount,
		DueAmount: primitives.DueAmount,
		Status:    primitives.Status,
	})

	return d, nil
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

func FromPrimitives(amount int, dueAmount int, status string) (*Document, error) {

	dueAmt, err := NewDueAmount(dueAmount)
	if err != nil {
		return nil, err
	}

	s, err := NewStatus(status)
	if err != nil {
		return nil, err
	}

	return new(amount, *dueAmt, *s)
}

func (d *Document) ToPrimitives() DocumentPrimitives {
	return DocumentPrimitives{
		Amount:    d.Amount,
		DueAmount: d.DueAmount.Value(),
		Status:    d.Status.Value(),
	}
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

func (d *Document) recordEvent(event any) {
	d.events = append(d.events, event)
}

func (d *Document) PullEvents() []any {
	events := d.events
	d.events = []any{}

	return events
}
