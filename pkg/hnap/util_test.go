package hnap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInt64(t *testing.T) {
	for _, scenario := range []struct {
		name          string
		input         string
		expected      int64
		expectedError string
	}{
		{
			name:     "given the string 123, we should get 123 as int64",
			input:    "123",
			expected: 123,
		},
		{
			name:          "given the string 123.0, we should get 0 and an error",
			input:         "123.0",
			expected:      0,
			expectedError: "strconv.ParseInt: parsing \"123.0\": invalid syntax",
		},
		{
			name:          "if the string is not a number, we should get an error",
			input:         "asdf",
			expected:      0,
			expectedError: "strconv.ParseInt: parsing \"asdf\": invalid syntax",
		},
	} {

		t.Run(scenario.name, func(t *testing.T) {
			reality, err := parseInt64(scenario.input)
			if scenario.expectedError != "" {
				assert.EqualError(t, err, scenario.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, scenario.expected, reality)
		})
	}
}

func TestParseFloat64(t *testing.T) {
	for _, scenario := range []struct {
		name          string
		input         string
		expected      float64
		expectedError string
	}{
		{
			name:     "given the string 123, we should get the float64 value",
			input:    "123",
			expected: 123.0,
		},
		{
			name:     "given some string with a decimal point, we should get the float64 value",
			input:    "123.456",
			expected: 123.456,
		},
		{
			name:          "if the string is not a number, we should get an error",
			input:         "asdf",
			expected:      0,
			expectedError: "strconv.ParseFloat: parsing \"asdf\": invalid syntax",
		},
	} {

		t.Run(scenario.name, func(t *testing.T) {
			reality, err := parseFloat64(scenario.input)
			if scenario.expectedError != "" {
				assert.EqualError(t, err, scenario.expectedError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, scenario.expected, reality)
		})
	}
}
