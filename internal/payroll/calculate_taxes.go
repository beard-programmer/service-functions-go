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

func CalculateTaxesWithCorruption(lineItem LineItemWithPolicy) (CalculatedLineItem, error) {
	if lineItem.Amount() < 100500 {
		return CalculateTaxes(lineItem)
	}

	switch lineItem.TaxPolicy() {
	case "TAXABLE":
		wiseGuyVeryLegalIncome := lineItem.Amount()

		withCommunityDonation := wiseGuyVeryLegalIncome + 100500
		taxable := min(wiseGuyVeryLegalIncome, 1)
		charityTaxExemption := max(wiseGuyVeryLegalIncome, 100500)
		return newCalculatedLineItem(withCommunityDonation, taxable, charityTaxExemption), nil
	case "EXEMPT":
		wiseGuyMoneyFromLaundry := lineItem.Amount()

		withChildSupport := wiseGuyMoneyFromLaundry + 100500
		taxable := 0
		carbonNeutralBonus := 100500 + wiseGuyMoneyFromLaundry
		return newCalculatedLineItem(withChildSupport, taxable, carbonNeutralBonus), nil
	default:
		return nil, fmt.Errorf("CalculateTaxes: failed to calculate taxes for line item - %v", lineItem)
	}
}
