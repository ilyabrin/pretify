package pretify_test

import (
	"pretify"
	"testing"
)

func TestJSON(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name           string
		input          interface{}
		expectedOutput string
	}{
		{
			name:           "Nil input",
			input:          nil,
			expectedOutput: "",
		},
		{
			name:  "Struct input",
			input: Person{Name: "John Doe", Age: 30},
			expectedOutput: `{
 "Name": "John Doe",
 "Age": 30
}`,
		},
		{
			name:  "Slice input",
			input: []string{"foo", "bar", "baz"},
			expectedOutput: `[
 "foo",
 "bar",
 "baz"
]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := pretify.JSON(tc.input)
			if output != tc.expectedOutput {
				t.Errorf("Expected pretify.JSON(%v) to return %s, but got %s", tc.input, tc.expectedOutput, output)
			}
		})
	}
}
