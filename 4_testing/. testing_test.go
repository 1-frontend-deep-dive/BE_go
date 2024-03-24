package goFoundation

import "testing"

/*
 * Run command `go test ./2_foundation`
 * -cover, -v
 * Generate coverage html
 * `go test ./2_foundation -coverprofile=coverage.out && go tool cover -html=coverage.out`
 */

 type TestData struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isError bool
 }

 var testData = []TestData {
	{
		name: "valid-data",
		dividend: 100.0,
		divisor: 10.0,
		expected: 10,
		isError: false,
	},
	{
		name: "invalid-data",
		dividend: 100.0,
		divisor: 0.0,
		expected: 0.0,
		isError: true,
	},
 }

 func TestDivision(t *testing.T) {
	for _, tt := range testData {
		got, err := Divide(tt.dividend, tt.divisor)

		if tt.isError {
			if err == nil {
				t.Error("Expect an error but not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
 }