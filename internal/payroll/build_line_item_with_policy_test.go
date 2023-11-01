package payroll

import "testing"

func TestBuildLineItemWithPolicy(t *testing.T) {
	testCases := []struct {
		name           string
		lineItem       LineItem
		expectedResult LineItemWithPolicy
		expectErr      bool
	}{
		{
			name:           "Taxable Line Item",
			lineItem:       NewLineItem(1000, "salary"),
			expectedResult: newLineItemWithPolicy(1000, "salary", "TAXABLE"),
			expectErr:      false,
		},
		{
			name:           "Taxable Line Item",
			lineItem:       NewLineItem(500, "bonus"),
			expectedResult: newLineItemWithPolicy(500, "bonus", "TAXABLE"),
			expectErr:      false,
		},
		{
			name:           "Exempt Line Item",
			lineItem:       NewLineItem(500, "meal_voucher"),
			expectedResult: newLineItemWithPolicy(500, "meal_voucher", "EXEMPT"),
			expectErr:      false,
		},
		{
			name:           "Unknown Line Item",
			lineItem:       NewLineItem(200, "unknown"),
			expectedResult: nil,
			expectErr:      true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Act: Call the function being tested
			result, err := BuildLineItemWithPolicy(testCase.lineItem)

			// Assert: Check if an error was expected
			if testCase.expectErr {
				// If an error was expected, make sure the function returned an error
				if err == nil {
					t.Fatalf("Expected an error but got nil")
				}

				// Make sure that no result was returned
				if result != nil {
					t.Fatal("Expected result to be nil when an error occurs, but got a result")
				}

				// Skip to the next iteration since we're done with this test case
				return
			}

			// If no error was expected, make sure the function did not return an error
			if err != nil {
				t.Fatalf("Did not expect an error but got: %v", err)
			}

			// Make sure a result was returned
			if result == nil {
				t.Fatal("Expected a result but got nil")
			}

			// Check if the result has the expected tax policy
			if result != testCase.expectedResult {
				t.Errorf("Expected TaxPolicy: %s, got: %s", testCase.expectedResult, result)
			}
		})
	}

}
