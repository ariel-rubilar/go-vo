package document_test

import (
	"testing"

	"github.com/ariel-rubilar/go-vo/document"
	"github.com/stretchr/testify/assert"
)

func TestDocument_UseCase(t *testing.T) {
	amount := 100

	doc := document.Document{
		Amount:    amount,
		DueAmount: amount,
		Status:    "pending",
	}

	assert.Equal(t, "pending", doc.Status)
	assert.Equal(t, 100, doc.Amount)
	assert.Equal(t, 100, doc.DueAmount)

}

type Request struct {
	Amount    int
	DueAmount int
	Status    string
}

func TestDocument_Handler(t *testing.T) {
	req := Request{
		Amount:    200,
		DueAmount: 100,
		Status:    "pending",
	}

	doc := document.Document{
		Amount:    req.Amount,
		DueAmount: req.DueAmount,
		Status:    "pending",
	}

	assert.Equal(t, "pending", doc.Status)
	assert.Equal(t, 200, doc.Amount)
	assert.Equal(t, 100, doc.DueAmount)
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

	doc := document.Document{
		Amount:    row.Amount,
		DueAmount: row.DueAmount,
		Status:    row.Status,
	}

	assert.Equal(t, "paid", doc.Status)
	assert.Equal(t, 300, doc.Amount)
	assert.Equal(t, 150, doc.DueAmount)
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

	doc := document.Document{
		Amount:    apiResp.Amount,
		DueAmount: apiResp.DueAmount,
		Status:    apiResp.Status,
	}

	assert.Equal(t, "overdue", doc.Status)
	assert.Equal(t, 400, doc.Amount)
	assert.Equal(t, 200, doc.DueAmount)
}
