package payroll

import (
	"reflect"
	"testing"
)

func TestCalculateEmployeePayroll(t *testing.T) {
	type args struct {
		buildLineItemWithPolicyFn BuildLineItemWithPolicyFn
		calculateTaxesFn          CalculateTaxesFn
		employeePayroll           EmployeePayroll
	}
	tests := []struct {
		name    string
		args    args
		want    CalculatedEmployeePayroll
		wantErr bool
	}{
		{
			name: "given correct employee payroll",
			args: args{
				BuildLineItemWithPolicy,
				CalculateTaxes,
				newEmployeePayroll(1, []LineItem{
					NewLineItem(100, "salary"),
					NewLineItem(150, "bonus"),
					NewLineItem(200, "meal_voucher"),
				}),
			},
			want:    newCalculatedEmployeePayroll(1, 450, 350, 100),
			wantErr: false,
		},
		{
			name: "given invalid employee payroll",
			args: args{
				BuildLineItemWithPolicy,
				CalculateTaxes,
				newEmployeePayroll(1, []LineItem{
					NewLineItem(100, "UNKNOWN_LINE-iTeM KEy12"),
					NewLineItem(150, "bonus"),
					NewLineItem(200, "meal_voucher"),
				}),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateEmployeePayroll(tt.args.buildLineItemWithPolicyFn, tt.args.calculateTaxesFn, tt.args.employeePayroll)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateEmployeePayroll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateEmployeePayroll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
