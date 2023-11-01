package payroll

import "fmt"

type CalculateTaxesFn func(LineItemWithPolicy) (CalculatedLineItem, error)

func CalculateTaxes(lineItem LineItemWithPolicy) (CalculatedLineItem, error) {
	switch lineItem.TaxPolicy() {
	case "TAXABLE":
		return newCalculatedLineItem(lineItem.Amount(), lineItem.Amount(), 0), nil
	case "EXEMPT":
		amount := lineItem.Amount()
		exempt := amount / 2
		taxable := amount - exempt
		return newCalculatedLineItem(amount, taxable, exempt), nil
	default:
		return nil, fmt.Errorf("CalculateTaxes: failed to calculate taxes for line item - %v", lineItem)
	}
}
