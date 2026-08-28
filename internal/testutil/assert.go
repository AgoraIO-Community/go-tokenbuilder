package testutil

import (
	"reflect"
)

type TestingT interface {
	Helper()
	Errorf(format string, args ...interface{})
}

func Nil(t TestingT, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Equal(t TestingT, expected, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v, got %#v", expected, actual)
	}
}
