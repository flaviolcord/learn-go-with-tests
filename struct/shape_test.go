package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got: %.2f, want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{10.0, 10.0}, 100.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{5, 3}, 7.2},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v, got: %g, want: %g", tt, got, tt.want)
			}
		})
	}
}
