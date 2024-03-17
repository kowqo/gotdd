package main

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, got, want float64) {
		t.Helper()

		if got != want {
			t.Errorf("want %f but got %f", want, got)
		}
	}

	t.Run("Base perimeter 10x10", func(t *testing.T) {
		r := Rectangle{
			width:  10.0,
			height: 10.0,
		}
		got := r.Perimeter()
		want := 40.0
		checkPerimeter(t, got, want)
	})

	t.Run("Base perimeter 5x10", func(t *testing.T) {
		r := Rectangle{
			width:  5.0,
			height: 10.0,
		}
		got := r.Perimeter()
		want := 30.0
		checkPerimeter(t, got, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12, 6}, 72.0},
		{"circle", Circle{10}, 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, want: 36},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("want %f but got %f", tt.want, got)
			}
		})

	}

}

func TestTriangle(t *testing.T) {

}
