package assert

import (
	"reflect"
	"testing"
)

func AssertEqual(want, got interface{}, t *testing.T) {
	isEqual := reflect.DeepEqual(want, got)
	if !isEqual {
		t.Errorf("assert equal failed:want %v got %v", want, got)
	}
}

func AssertNotEqual(want, got interface{}, t *testing.T) {
	isEqual := reflect.DeepEqual(want, got)
	if isEqual {
		t.Errorf("assert not equal failed::want %v got %v", want, got)
	}
}
