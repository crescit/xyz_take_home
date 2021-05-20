# TrueAccord_take_home
Hello TrueAccord, Thank you for going through the process with me. Here's the private repository for the take home assignment. 

To test the application with other urls replace the corresponding values in the .env
  
  DEBT_URL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts"
	PAYMENT_PLAN_URL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans"
	PAYMENT_URL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments"

Running the application:
  To run the application via docker:
    docker build -t trueaccordtakehome .
    docker run -d trueaccordtakehome  (this outputs <container id>)
    docker logs <container id>

  To run the application locally (further build support can be found via go's documentation):
    go run main.go

  To run tests locally: 
    go test *.go
  
Steps for improvement:
  The urls for the endpoints are pretty much hardcoded.
  Float64 isn't safe to use on money, I would ideally convert to integer or use a third party library instead of rounding off the float.
  I would make main.go less verbose and continue to bring things out into their own functions as I did the API fetching logic.
  The tests currently just test the output and the data ingestion, tests could be written to thoroughly test the data. 
  Edge cases that would've been great to handle: negative payment amounts, payment amounts greater than debt, improperly formatted dates.
  Payments seem to only be associated with a payment plan, it doesn't seem possible to pay off a debt that doesn't have a payment plan currently. 
  
Steps the code does:
  Fetches the data for debts, payment plans, and payments.
  
  Initializes three maps:
    Map 1: debtToPlan, maps the debts id to a payment plan
    Map 2: payToPlan, maps the payment plan id to the total amount of payments accrued for the plan
    Map 3: dateToPlan, maps the payment plan id to the latest date of a payment on the plan
   
  For each debt:
    Calculate is_in_payment_plan by checking that debt id exists in map 1 and that amount to pay is less than debt.
    
    Calculate remaining_amount 
      If the debt is associated with a payment plan subtract payments from payment plan total else it sets it to the original debt amount.
      
    Calculate next_payment_due_date
      Start at the payment plans start_date, until it is greater than the latest payment date continue adding it's corresponding increment. 
      
  Print Debts In JSON
    For each debt in the response, print out the struct in JSON in the following format with a \n at the end. 
    {
      "id": 3,
      "amount": 12938,
      "is_in_payment_plan": true,
      "remaining_amount": 622.41,
      "next_payment_due_date": "2020-08-15"
    }
