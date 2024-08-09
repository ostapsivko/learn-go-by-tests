package shapes

import (
	"math"
	"testing"
)

func TestPerimiter(t *testing.T) {
	rect := Rectangle{2.5, 2.5}
	got := rect.Perimeter()
	want := 10.0

	if want != got {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	t.Run("test area of a rectangle", func(t *testing.T) {
		rect := Rectangle{2.5, 2.5}
		got := rect.Area()
		want := 6.25

		if want != got {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})

	t.Run("test area of a rectangle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := math.Pi * 100

		if want != got {
			t.Errorf("want %g, got %g", want, got)
		}
	})
}
