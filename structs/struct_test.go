package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	// when
	result := Perimeter(Rectangle{10.0, 10.0})

	// then
	assert.Equal(t, 40.0, result)
}

func TestArea(t *testing.T) {
	tests := []struct { // this syntax is insane and not in a good way
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedArea, test.shape.Area())
		})
	}
}
