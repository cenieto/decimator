package geom3d

import (
	"decimator/pkg/testutils"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestNewVector validates the creation of a Vector3D using NewVector.
//
// This test checks if the Vector3D created by NewVector matches the expected
// 3D vector representation.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestNewVector(t *testing.T) {
	input := []float64{0.0, 1.0, 2.0}
	result := NewVector(input[0], input[1], input[2])
	expected := mat.NewVecDense(3, input)

	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("NewVector(%v) = %v; want %v", input, mat.Formatted(result), mat.Formatted(expected))
	}
}

// TestNewVectorTwoPoints validates the creation of a vector using two points.
//
// This test checks if the vector created by NewVectorTwoPoints matches the
// expected 3D vector representation when calculated as Point2 - Point1.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestNewVectorTwoPoints(t *testing.T) {
	p1 := NewPoint(0.0, 1.0, 2.0)
	p2 := NewPoint(1.0, 0.0, 2.0)
	result := NewVectorTwoPoints(p1, p2)

	expected := mat.NewVecDense(3, []float64{1.0, -1.0, 0.0})

	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("NewVectorTwoPoints(%v, %v) = %v; want %v", mat.Formatted(p1), mat.Formatted(p2), mat.Formatted(result), mat.Formatted(expected))
	}
}
