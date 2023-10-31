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

func NewCalculatedLineItem(amount int, taxable int, exempt int) CalculatedLineItem {
	return calculatedLineItem{amount, taxable, exempt}
}

type LineItem interface {
	Amount() int
	LineItemKey() string
}

type LineItemWithPolicy interface {
	LineItem
	TaxPolicy() string
}

type CalculatedLineItem interface {
	Amount() int
	Taxable() int
	Exempt() int
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

func (li calculatedLineItem) Amount() int {
	return li.amount
}
func (li calculatedLineItem) Taxable() int {
	return li.taxable
}
func (li calculatedLineItem) Exempt() int {
	return li.exempt
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

type calculatedLineItem struct {
	amount  int
	taxable int
	exempt  int
}
