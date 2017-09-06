package errchk

import (
	"errors"
	"testing"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestCheck(t *testing.T) {
	cases := []struct {
		err error
		m   string
	}{
		{nil, "no error"},
		{errors.New("some error"), "some message"},
	}
	for _, tt := range cases {
		if tt.err != nil {
			assertPanic(t, func() { Check(tt.err, tt.m) })
		} else {
			Check(tt.err, tt.m)
		}
	}
}
