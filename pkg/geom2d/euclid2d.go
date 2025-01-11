package geom2d

import (
	"decimator/pkg/geom3d"
)

// Euclid2D represents a 2D geometric system.
// It provides the necessary methods to perform 2D geometric operations such as cross product and distance calculations.
type Euclid2D struct {
}

// NewEuclid creates and returns a new instance of Euclid2D.
//
// Returns:
//   - Euclid2D: A new instance of the 2D geometry system.
func NewEuclid() Euclid2D {
	return Euclid2D{}
}

// Dimension returns the dimension of the geometry system.
//
// Returns:
//   - int: The dimension of the geometry, which is always 2 for this system.
func (g Euclid2D) Dimension() int {
	return 2
}

// CrossProduct computes the cross product of two 2D vectors and returns the result as a 3D vector.
// The Z-component of the resulting 3D vector represents the scalar cross product of the 2D vectors.
//
// Parameters:
//   - v1 (*Vector2D): The first 2D vector to be used in the cross product.
//   - v2 (*Vector2D): The second 2D vector to be used in the cross product.
//
// Returns:
//   - *geom3d.Vector3D: A 3D vector where the Z-component represents the cross product of the input 2D vectors.
func (g Euclid2D) CrossProduct(v1, v2 *Vector2D) *geom3d.Vector3D {
	result := geom3d.NewVector(
		0.0,
		0.0,
		v1.At(0, 0)*v2.At(1, 0)-v1.At(1, 0)*v2.At(0, 0),
	)
	return result
}

// CrossProductNorm computes the magnitude (norm) of the cross product of two 2D vectors.
// This magnitude corresponds to the area of the parallelogram formed by the vectors.
//
// Parameters:
//   - v1 (*Vector2D): The first 2D vector.
//   - v2 (*Vector2D): The second 2D vector.
//
// Returns:
//   - float64: The magnitude (norm) of the cross product.
func (g Euclid2D) CrossProductNorm(v1, v2 *Vector2D) float64 {
	crossProduct := g.CrossProduct(v1, v2)
	result := crossProduct.VecDense.Norm(2)
	return result
}

// DoubleAreaTriangle calculates the double of the area of a triangle formed by a point and a line.
// The area is computed using the cross product of the vector from the point to one line endpoint
// and the line's direction vector.
//
// Parameters:
//   - point (*Point2D): The point used to form the triangle.
//   - line (*Line2D): The line forming the base of the triangle.
//
// Returns:
//   - float64: The double of the triangle's area.
func (g Euclid2D) DoubleAreaTriangle(point *Point2D, line *Line2D) float64 {
	lineToPoint := NewVectorTwoPoints(point, line.Point1)
	vectorDirector := line.VectorDirector()
	numerator := g.CrossProductNorm(lineToPoint, vectorDirector)
	return numerator
}

// DistancePointLine computes the shortest distance from a point to a line.
// It divides the double area of the triangle formed by the point and the line
// by the norm of the line's direction vector.
//
// Parameters:
//   - point (*Point2D): The point whose distance to the line is being calculated.
//   - line (*Line2D): The line to which the distance is being measured.
//
// Returns:
//   - float64: The shortest distance from the point to the line.
func (g Euclid2D) DistancePointLine(point *Point2D, line *Line2D) float64 {
	numerator := g.DoubleAreaTriangle(point, line)
	denominator := line.VectorDirector().Norm(2)
	return numerator / denominator
}
