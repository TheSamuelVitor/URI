package assert

import "testing"

func Equal(t *testing.T, val1, val2 interface{}) {
	if val1 != val2 {
		t.Errorf("\ngot: %v\nwant: %v", val1, val2)
	}
}
