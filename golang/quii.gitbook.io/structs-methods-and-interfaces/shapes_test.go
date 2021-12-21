package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectange := Rectange{10.0, 10.0}
	got := rectange.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectange{12.0, 6.0}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
