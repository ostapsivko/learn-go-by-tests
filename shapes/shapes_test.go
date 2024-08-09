package shapes

import "testing"

func TestPerimiter(t *testing.T) {
	got := Perimeter(2.5, 2.5)
	want := 10.0

	if want != got {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(2.5, 2.5)
	want := 6.25

	if want != got {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}
