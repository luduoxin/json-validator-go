package validator

import "testing"

var validTests = []struct {
	data string
	ok   bool
}{
	{`foo`, false},
	{`}{`, false},
	{`{]`, false},
	{`123`, false},
	{`123.5`, false},
	{`0.5`, false},
	{`true`, false},
	{`false`, false},
	{`null`, false},
	{`[1,2,3]`, false},
	{`["a","b","c"]`, false},
	{`{"foo":bar}`, false},
	{`[{"foo":"bar"},]`, false},
	{`{}`, true},
	{`[]`, true},
	{`[{}]`, true},
	{`{"foo":"bar"}`, true},
	{`{"foo":"bar","bar":{"baz":["qux"]}}`, true},
	{`[{"a":"[\"c\":\"d\"]"}]`, true},
	{`[{"a":[]}]`, true},
	{` {"a":"b"}`, true},
	{` {"a":"b"} `, true},
	{`{"a":"b"} `, true},
	{`{"a": "b"} `, true},
}

func TestValid(t *testing.T) {
	for _, tt := range validTests {
		if ok := Valid([]byte(tt.data)); ok != tt.ok {
			t.Errorf("Valid(%#q) = %v, want %v", tt.data, ok, tt.ok)
		}
	}
}
