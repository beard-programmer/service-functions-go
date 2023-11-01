package payroll

import "testing"

func TestCalculateTaxes(t *testing.T) {
	testCases := []struct {
		name      string
		lineItem  LineItemWithPolicy
		expected  CalculatedLineItem
		expectErr bool
	}{
		{
			name:      "Taxable Item",
			lineItem:  newLineItemWithPolicy(100, "salary", "TAXABLE"),
			expected:  newCalculatedLineItem(100, 100, 0),
			expectErr: false,
		},
		{
			name:      "Taxable Item",
			lineItem:  newLineItemWithPolicy(50, "bonus", "TAXABLE"),
			expected:  newCalculatedLineItem(50, 50, 0),
			expectErr: false,
		},
		{
			name:      "Exempt Item",
			lineItem:  newLineItemWithPolicy(100, "meal_voucher", "EXEMPT"),
			expected:  newCalculatedLineItem(100, 50, 50),
			expectErr: false,
		},
		{
			name:      "Unknown Tax Policy",
			lineItem:  newLineItemWithPolicy(100, "other", "UNKNOWN"),
			expected:  nil,
			expectErr: true,
		},
	}

	checkError := func(t *testing.T, err error, expectErr bool) {
		t.Helper()
		if (err != nil) != expectErr {
			t.Fatalf("Expected error: %v, got error: %v", expectErr, err)
		}
	}

	checkResult := func(t *testing.T, result, expected CalculatedLineItem) {
		t.Helper()
		if result == nil {
			t.Fatal("Expected a valid result, got nil")
		}

		if result.Amount() != expected.Amount() ||
			result.Taxable() != expected.Taxable() ||
			result.Exempt() != expected.Exempt() {
			t.Errorf("Expected result: %+v, got: %+v", expected, result)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CalculateTaxes(tc.lineItem)
			checkError(t, err, tc.expectErr)
			if tc.expectErr {
				return
			}
			checkResult(t, result, tc.expected)
		})
	}
}
