package main

import (
	"testing"

	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
)

var (
	cubeTest              Cube    = Cube{4}
	expectedVolume        float64 = 64.00
	expectedArea          float64 = 96.00
	expectedCircumference float64 = 48.00
)

func TestVolumeCalculate(t *testing.T) {
	t.Logf("Volume: %2.f", cubeTest.CubeVolume())

	if cubeTest.CubeVolume() != expectedVolume {
		t.Errorf("WRONG, expected volume: %2.f", expectedVolume)
	}
}

func TestCircumferenceCalculate(t *testing.T) {
	t.Logf("Circumference: %2.f", cubeTest.CubeCircumference())

	if cubeTest.CubeCircumference() != expectedCircumference {
		t.Errorf("WRONG, expected circumference: %.2f", expectedCircumference)
	}
}

func TestAreaCalculate(t *testing.T) {
	t.Logf("Volume: %.2f", cubeTest.CubeArea())

	if cubeTest.CubeArea() != expectedArea {
		t.Errorf("WRONG, expected area: %.2f", expectedArea)
	}
}

func BenchmarkCalculateArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeTest.CubeArea()
	}
}

// Using testing library -> Testify
func TestVolumeCalculateTestify(t *testing.T) {
	assert.Equal(t, cubeTest.CubeVolume(), expectedVolume, "WRONG Volume calculate")
}

func TestAreaCalculateTestify(t *testing.T) {
	assert.Equal(t, cubeTest.CubeArea(), expectedArea, "WRONG Area calculate")
}

func TestCircumferenceCalculateTestify(t *testing.T) {
	assert.Equal(t, cubeTest.CubeCircumference(), expectedCircumference, "WRONG Circumference calculate")
}
