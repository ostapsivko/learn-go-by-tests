package shapes

import (
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

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rect", shape: Rectangle{12, 6}, want: 72.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g, want %g", tt.shape, got, tt.want)
			}
		})
	}
}
