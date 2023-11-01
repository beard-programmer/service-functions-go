package payroll

import "fmt"

type BuildLineItemWithPolicyFn func(LineItem) (LineItemWithPolicy, error)

func BuildLineItemWithPolicy(lineItem LineItem) (LineItemWithPolicy, error) {
	switch lineItem.LineItemKey() {
	case "salary", "bonus":
		return newLineItemWithPolicy(lineItem.Amount(), lineItem.LineItemKey(), "TAXABLE"), nil
	case "meal_voucher":
		return newLineItemWithPolicy(lineItem.Amount(), lineItem.LineItemKey(), "EXEMPT"), nil
	default:
		return nil, fmt.Errorf("BuildLineItemWithPolicy: failed to find proper tax polciy for line item - %v", lineItem)
	}
}
