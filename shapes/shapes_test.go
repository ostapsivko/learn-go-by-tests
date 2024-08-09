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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want %f, got %f", want, got)
		}
	}

	t.Run("test area of a rectangle", func(t *testing.T) {
		rect := Rectangle{2.5, 2.5}
		checkArea(t, rect, 6.25)
	})

	t.Run("test area of a circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, math.Pi*100)
	})
}
