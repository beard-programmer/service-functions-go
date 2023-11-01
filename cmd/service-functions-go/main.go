package main

import (
	"fmt"
	"service-functions-go/internal/payroll"
)

func main() {
	fmt.Println("Hello main")

	employeePayroll := payroll.NewEmployeePayroll(1, []payroll.LineItem{
		payroll.NewLineItem(100, "salary"),
		payroll.NewLineItem(150, "bonus"),
		payroll.NewLineItem(200, "meal_voucher"),
	})
	result, err := payroll.CalculateEmployeePayroll(payroll.BuildLineItemWithPolicy, payroll.CalculateTaxesWithCorruption, employeePayroll)
	if err != nil {
		fmt.Printf("Failed to calculate, reason: %v\n", err)
	} else {
		fmt.Printf("Calculated ep: %v and got result: %v\n", employeePayroll, result)
	}
}
