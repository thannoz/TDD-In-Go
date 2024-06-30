package arrayslices

import (
	"reflect"
	"testing"
)

/*
1. Deklariere und initialisiere ein Array von 5 ganzen Zahlen und fülle es mit den Werten 1 bis 5.

Array-Elemente ändern
2. Ändere das zweite Element des Arrays in den Wert 10 und drucke das Array aus.

Array-Länge
3. Schreibe eine Funktion, die die Länge eines Arrays zurückgibt. */

func TestArrayToIncludeFiveDigits(t *testing.T) {
	t.Parallel()

	expected := [5]int{1, 2, 3, 4, 5}

	actual := initializeArray()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
