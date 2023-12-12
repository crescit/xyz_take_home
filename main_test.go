package main

import (
	"testing"
)

// tests if function prints an appropriate jsonl response on correct input
func TestPrintInJson(t *testing.T) {
	var testDebt1 Debt
	testDebt1.ID = 2
	testDebt1.Amount = 123.58
	var testDebt2 Debt
	testDebt2.ID = 2
	testDebt2.Amount = 123.58
	var d = []Debt{testDebt1, testDebt2}

	expected := "{\n  \"id\": 2,\n  \"amount\": 123.58,\n  \"is_in_payment_plan\": false,\n  \"remaining_amount\": 0,\n  \"next_payment_due_date\": \"\"\n}\n{\n  \"id\": 2,\n  \"amount\": 123.58,\n  \"is_in_payment_plan\": false,\n  \"remaining_amount\": 0,\n  \"next_payment_due_date\": \"\"\n}\n"
	result := PrintInJson(d)
	if expected != result {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}

// tests for error in case of empty debt url
func TestGetDebt(t *testing.T) {
	var testUrl = ""
	_, err := GetDebt(testUrl)
	if err == nil {
		t.Errorf("function failed to fail on empty response with url %q", testUrl)
	}

	testUrl = ""
	debt, err := GetDebt(testUrl)

	if len(debt) != 5 {
		t.Errorf("function failed to get appropriate response with url %q", testUrl)
	}
}

func TestGetPaymentPlans(t *testing.T) {
	var testUrl = ""
	_, err := GetPaymentPlans(testUrl)
	if err == nil {
		t.Errorf("function failed to fail on empty response with url %q", testUrl)
	}

	testUrl = ""
	debt, err := GetPaymentPlans(testUrl)

	if len(debt) != 4 {
		t.Errorf("function failed to get appropriate response with url %q", testUrl)
	}
}

func TestGetPayments(t *testing.T) {
	var testUrl = ""
	_, err := GetPayments(testUrl)
	if err == nil {
		t.Errorf("function failed to fail on empty response with url %q", testUrl)
	}

	testUrl = ""
	debt, err := GetPayments(testUrl)

	if len(debt) != 8 {
		t.Errorf("function failed to get appropriate response with url %q", testUrl)
	}
}
