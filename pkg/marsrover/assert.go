package marsrover

import (
	"reflect"
	"testing"
)

func assertEqual(want, got interface{}, t *testing.T) {
	isEqual := reflect.DeepEqual(want, got)
	if !isEqual {
		t.Errorf("assert equal failed:want %v got %v", want, got)
	}
}

func assertNotEqual(want, got interface{}, t *testing.T) {
	isEqual := reflect.DeepEqual(want, got)
	if isEqual {
		t.Errorf("assert not equal failed::want %v got %v", want, got)
	}
}
