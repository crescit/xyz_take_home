package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
)

// program fetches data from 3 endpoints, parses the data and prints it out in JSON
func main() {
	format := "2006-01-02"
	debts := getDebt("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts")
	paymentPlans := getPaymentPlans("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans")
	payments := getPayments("https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments")

	// init maps of payments and payment plans for quick lookup
	debtToPlan := make(map[int]int)       // debt_id to index in paymentPlans[]
	payToPlan := make(map[int]float64)    // payment_plan_id to total amount paid
	dateToPlan := make(map[int]time.Time) // payment_plan_id to last date a payment occured

	for i := 0; i < len(paymentPlans); i += 1 {
		debtToPlan[paymentPlans[i].DebtID] = i
	}
	for i := 0; i < len(payments); i += 1 {
		payToPlan[payments[i].ID] += payments[i].Amount
		date, _ := time.Parse(format, payments[i].Date)
		if dateToPlan[payments[i].ID].Before(date) {
			dateToPlan[payments[i].ID] = date
		}
	}

	// add is_in_payment_plan, remaining_amount, and next_payment_due_date to debt rows
	for i := 0; i < len(debts); i += 1 {
		// calculate is_in_payment_plan, by seeing that id exists in map and that amount to pay is < debt
		paymentPlanIndex := debtToPlan[debts[i].ID]
		pid := paymentPlans[paymentPlanIndex].ID
		if paymentPlans[paymentPlanIndex].DebtID == debts[i].ID && payToPlan[pid] < debts[i].Amount {
			debts[i].IsInPaymentPlan = true
		}

		// calculate remaining_amount, if the debt is associated with a payment plan subtract payments from payment plan total else set to debt amount
		if paymentPlans[debtToPlan[debts[i].ID]].DebtID == debts[i].ID {
			debts[i].RemainingAmount = math.Floor((paymentPlans[debtToPlan[debts[i].ID]].AmountToPay-payToPlan[paymentPlans[debtToPlan[debts[i].ID]].ID])*100) / 100
		} else {
			debts[i].RemainingAmount = debts[i].Amount
		}
		if debts[i].RemainingAmount == 0 {
			debts[i].IsInPaymentPlan = false
		}

		// calculate "next_payment_due_date",
		if debts[i].RemainingAmount != 0 && debts[i].IsInPaymentPlan {
			pid := debtToPlan[debts[i].ID]
			freq := paymentPlans[pid].Frequency
			start, _ := time.Parse(format, paymentPlans[pid].StartDate)
			for ok := true; ok; ok = start.Before(dateToPlan[pid]) {
				var weeksToAdd int
				if freq == "WEEKLY" {
					weeksToAdd = 1
				}
				if freq == "BI_WEEKLY" {
					weeksToAdd = 2
				}
				start = start.Add(time.Hour * 24 * 7 * time.Duration(weeksToAdd))
			}
			debts[i].NextPayment = start.Format(format)
		}
	}

	PrintInJson(debts)
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

func PrintInJson(debts []Debt) string {
	var res string
	for i := 0; i < len(debts); i += 1 {
		debt, err := json.MarshalIndent(debts[i], "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(debt) + "\n")
		res += string(debt) + "\n"
	}
	return res
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
