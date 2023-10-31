package payroll

func NewLineItem(amount int, lineItemKey string) LineItem {
	return lineItem{amount, lineItemKey}
}

func NewLineItemWithPolicy(amount int, lineItemKey string, taxPolicy string) LineItemWithPolicy {
	return lineItemWithPolicy{
		amount,
		lineItemKey,
		taxPolicy,
	}
}

type LineItem interface {
	Amount() int
	LineItemKey() string
}

type LineItemWithPolicy interface {
	LineItem
	TaxPolicy() string
}

func (li lineItem) Amount() int {
	return li.amount
}

func (li lineItem) LineItemKey() string {
	return li.lineItemKey
}

func (li lineItemWithPolicy) Amount() int {
	return li.amount
}

func (li lineItemWithPolicy) LineItemKey() string {
	return li.lineItemKey
}

func (li lineItemWithPolicy) TaxPolicy() string {
	return li.taxPolicy
}

type lineItem struct {
	amount      int
	lineItemKey string
}

type lineItemWithPolicy struct {
	amount      int
	lineItemKey string
	taxPolicy   string
}
