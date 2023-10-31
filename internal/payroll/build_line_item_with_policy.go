package payroll

import "fmt"

type BuildLineItemWithPolicyFn func(LineItem) (LineItemWithPolicy, error)

func BuildLineItemWithPolicy(lineItem LineItem) (LineItemWithPolicy, error) {
	switch lineItem.LineItemKey() {
	case "salary", "bonus":
		return NewLineItemWithPolicy(lineItem, "TAXABLE"), nil
	case "meal_voucher":
		return NewLineItemWithPolicy(lineItem, "EXEMPT"), nil
	default:
		return nil, fmt.Errorf("BuildLineItemWithPolicy: failed to find proper tax polciy for line item - %v", lineItem)
	}
}
