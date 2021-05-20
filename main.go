package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/**
* Debt consumes the response from:
*		https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts
*
*		id (integer)
*		amount (real) - amount owed in USD
*
* Payments Plans consumes the response from:
* 	https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans
*
*    id (integer)
*    debt_id (integer) - The associated debt.
*    amount_to_pay (real) - Total amount (in USD) needed to be paid to resolve this payment plan.
*    installment_frequency (text) - The frequency of payments. Is one of: WEEKLY or BI_WEEKLY (14 days).
*    installment_amount (real) - The amount (in USD) of each payment installment.
*    start_date (string) - ISO 8601 date of when the first payment is due.
*
*
*	Payments consumes the response from:
*		https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments
*
*   payment_plan_id (integer)
*   amount (real)
*   date (string) - ISO 8601 date of when this payment occurred.
*
 */

// call api
// parse response to array
// print valid line of json
// write tests
// zip it
// be done
// https://gist.github.com/jeffling/2dd661ff8398726883cff09839dc316c

func main() {
	debts := getDebt("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts")
	paymentPlans := getPaymentPlans("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans")
	payments := getPayments("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments")

	// init maps of payments and payment plans for quick lookup
	debtToPlan := make(map[int]int)    // debt_id to index in paymentPlans[]
	payToPlan := make(map[int]float64) // payment_plan_id to total amount paid
	for i := 0; i < len(paymentPlans); i += 1 {
		debtToPlan[paymentPlans[i].DebtID] = i
	}
	for i := 0; i < len(payments); i += 1 {
		payToPlan[payments[i].ID] += payments[i].Amount
	}

	for i := 0; i < len(debts); i += 1 {
		// calculate is_in_payment_plan, by seeing that id exists in map and that amount to pay is < debt
		paymentPlanIndex := debtToPlan[debts[i].ID]
		pid := paymentPlans[paymentPlanIndex].ID
		if paymentPlans[paymentPlanIndex].DebtID == debts[i].ID && payToPlan[pid] < debts[i].Amount {
			debts[i].IsInPaymentPlan = true
		}

		// calculate remaining_amount, if the debt is associated with a payment plan subtract payments from payment plan total else set to debt amoun
		if paymentPlans[debtToPlan[debts[i].ID]].DebtID == debts[i].ID {
			debts[i].RemainingAmount = paymentPlans[debtToPlan[debts[i].ID]].AmountToPay - payToPlan[paymentPlans[debtToPlan[debts[i].ID]].ID]
		} else {
			debts[i].RemainingAmount = debts[i].Amount
		}

		// date, _ := time.Parse("2006-01-02", paymentPlans[pid].StartDate)
		// log.Printf("%v %v %v  \n", paymentPlanIndex, payToPlan[pid], debts[i])
		log.Printf("%v", debts[i])
	}
}

func getDebt(url string) []Debt {
	res, _ := http.Get(url)
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var data []Debt
	decoder.Decode(&data)
	return data
}

func getPaymentPlans(url string) []PaymentPlan {
	res, _ := http.Get(url)
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var data []PaymentPlan
	decoder.Decode(&data)
	return data
}

func getPayments(url string) []Payment {
	res, _ := http.Get(url)
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var data []Payment
	decoder.Decode(&data)
	return data
}

type Debt struct {
	ID              int     `form:"id" json:"id"`
	Amount          float64 `form:"amount" json:"amount"`
	IsInPaymentPlan bool    `form:"is_in_payment_plan" json:"is_in_payment_plan"`
	RemainingAmount float64 `form:"remaining_amount" json:"remaining_amount"`
	NextPayment     string  `form:"next_payment_due_date" json:"next_payment_due_date"`
}

type PaymentPlan struct {
	ID                int     `form:"id" json:"id"`
	DebtID            int     `form:"debt_id" json:"debt_id"`
	AmountToPay       float64 `form:"amount_to_pay" json:"amount_to_pay"`
	Frequency         string  `form:"installment_frequency" json:"installment_frequency"`
	InstallmentAmount float64 `form:"installment_amount" json:"installment_amount"`
	StartDate         string  `form:"start_date" json:"start_date"`
}

type Payment struct {
	ID     int     `form:"payment_plan_id" json:"payment_plan_id"`
	Amount float64 `form:"amount" json:"amount"`
	Date   string  `form:"date" json:"date"`
}
