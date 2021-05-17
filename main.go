package main

import (
	"log"
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

func main() {
	log.Print("executed")
}
