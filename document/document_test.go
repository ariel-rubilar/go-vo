package document_test

import (
	"testing"

	"github.com/ariel-rubilar/go-vo/document"
	"github.com/stretchr/testify/assert"
)

func TestDocument_UseCase(t *testing.T) {
	amount := 100

	doc, err := document.Create(amount)

	assert.NoError(t, err)

	doc.Status.IsPaid()

	assert.Equal(t, "pending", doc.Status.Value())
	assert.Equal(t, 100, doc.Amount)
	assert.Equal(t, 100, doc.DueAmount.Value())

}

type Request struct {
	Amount    int
	Status    string
	DueAmount int
}

func TestDocument_Handler(t *testing.T) {
	req := Request{
		Amount:    100,
		Status:    "pending",
		DueAmount: 100,
	}

	doc, err := document.FromPrimitives(req.Amount, req.DueAmount, req.Status)

	assert.NoError(t, err)

	assert.Equal(t, "pending", doc.Status.Value())
	assert.Equal(t, 100, doc.Amount)
	assert.Equal(t, 100, doc.DueAmount.Value())
}

type Row struct {
	Amount    int
	DueAmount int
	Status    string
}

func TestDocument_FromDB(t *testing.T) {

	row := Row{
		Amount:    300,
		DueAmount: 150,
		Status:    "paid",
	}

	doc := document.Rehydrate(row.Amount, row.DueAmount, row.Status)

	assert.Equal(t, "paid", doc.Status.Value())
	assert.Equal(t, 300, doc.Amount)
	assert.Equal(t, 150, doc.DueAmount.Value())
}

type ApiResponse struct {
	Amount    int
	DueAmount int
	Status    string
}

func TestDocument_FromApi(t *testing.T) {

	apiResp := ApiResponse{
		Amount:    400,
		DueAmount: 200,
		Status:    "overdue",
	}

	doc, err := document.FromPrimitives(apiResp.Amount, apiResp.DueAmount, apiResp.Status)

	assert.NoError(t, err)

	assert.Equal(t, "overdue", doc.Status.Value())
	assert.Equal(t, 400, doc.Amount)
	assert.Equal(t, 200, doc.DueAmount.Value())
}
