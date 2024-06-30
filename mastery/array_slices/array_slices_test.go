package arrayslices

import (
	"reflect"
	"testing"
)

func TestArrayToIncludeFiveDigits(t *testing.T) {
	t.Parallel()

	expected := [5]int{1, 2, 3, 4, 5}

	actual := initializeArray()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
