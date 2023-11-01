package payroll

func newEmployeePayroll(employeeId int, lineItems []LineItem) EmployeePayroll {
	return employeePayroll{employeeId, lineItems}
}

func newCalculatedEmployeePayroll(employeeId int, total int, taxable int, exempt int) CalculatedEmployeePayroll {
	return calculatedEmployeePayroll{employeeId, total, taxable, exempt}
}

type EmployeePayroll interface {
	EmployeeId() int
	LineItems() []LineItem
}

type CalculatedEmployeePayroll interface {
	EmployeeId() int
	Total() int
	Taxable() int
	Exempt() int
}

func (e employeePayroll) EmployeeId() int {
	return e.employeeId
}

func (e employeePayroll) LineItems() []LineItem {
	return e.lineItems
}

func (c calculatedEmployeePayroll) EmployeeId() int {
	return c.employeeId
}

func (c calculatedEmployeePayroll) Total() int {
	return c.total
}

func (c calculatedEmployeePayroll) Taxable() int {
	return c.taxable
}

func (c calculatedEmployeePayroll) Exempt() int {
	return c.exempt
}

type employeePayroll struct {
	employeeId int
	lineItems  []LineItem
}

type calculatedEmployeePayroll struct {
	employeeId int
	total      int
	taxable    int
	exempt     int
}
