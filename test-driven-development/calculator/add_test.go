package calculator

import "testing"

func TestAddPositiveNumbers(t *testing.T) {
	expected := 9
	actual := Add(4, 5)
	if expected != actual {
		t.Errorf("add(4, 5) should be %d, get %d", expected, actual)
	}
}

func TestAddNegativeNumbers(t *testing.T) {
	expected := -1
	actual := Add(4, -5)
	if expected != actual {
		t.Errorf("add(4, -5) should be %d, get %d", expected, actual)
	}
}
