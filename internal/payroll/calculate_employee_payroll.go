package payroll

import "fmt"

type CalculateEmployeePayrollFn func(BuildLineItemWithPolicyFn, CalculateTaxesFn, EmployeePayroll) (CalculatedEmployeePayroll, error)

func CalculateEmployeePayroll(
	buildLineItemWithPolicyFn BuildLineItemWithPolicyFn, // Dependency
	calculateTaxesFn CalculateTaxesFn, // Dependency
	employeePayroll EmployeePayroll,
) (CalculatedEmployeePayroll, error) {
	// Step 1, 2: receive line items and build line items with policy using Dependency
	var lineItems []LineItemWithPolicy
	for _, li := range employeePayroll.LineItems() {
		result, err := buildLineItemWithPolicyFn(li)
		if err != nil {
			return nil, fmt.Errorf("CalculateEmployeePayroll: failed to calculate employees payroll %v. %v", employeePayroll, err)
		}
		lineItems = append(lineItems, result)
	}

	// Step 3: calculate taxes for each line item using Dependency
	var calculatedLineItems []CalculatedLineItem
	for _, li := range lineItems {
		result, err := calculateTaxesFn(li)
		if err != nil {
			return nil, fmt.Errorf("CalculateEmployeePayroll: failed to calculate employees payroll %v. %v", employeePayroll, err)
		}
		calculatedLineItems = append(calculatedLineItems, result)
	}

	// Step 4: build CalculatedEmployeePayroll
	total := 0
	taxable := 0
	exempt := 0
	for _, li := range calculatedLineItems {
		total += li.Amount()
		taxable += li.Taxable()
		exempt += li.Exempt()
	}

	return newCalculatedEmployeePayroll(employeePayroll.EmployeeId(), total, taxable, exempt), nil
}
