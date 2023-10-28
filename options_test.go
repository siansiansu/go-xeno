package xeno

import (
	"net/url"
	"reflect"
	"testing"
)

func TestProcessOptions(t *testing.T) {
	tests := []struct {
		name     string
		options  []RequestOption
		expected url.Values
	}{
		{
			name: "Page option sets the correct page value",
			options: []RequestOption{
				Page(1),
			},
			expected: url.Values{"page": []string{"1"}},
		},
		{
			name:     "No options provided",
			options:  []RequestOption{},
			expected: url.Values{},
		},
		{
			name: "Multiple options set correct values",
			options: []RequestOption{
				Page(2),
				func(o *requestOptions) {
					o.urlParams.Set("size", "10")
				},
			},
			expected: url.Values{"page": []string{"2"}, "size": []string{"10"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := processOptions(test.options...)
			if !reflect.DeepEqual(result.urlParams, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result.urlParams)
			}
		})
	}
}
