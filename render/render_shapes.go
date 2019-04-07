package render

import (
	"image/color"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/gl"
)

// TriangleType is the type of triangle: Right or Isosceles.
type TriangleType uint8

const (
	// TriangleIsosceles indicates a Triangle where two sides have equal length
	TriangleIsosceles TriangleType = iota
	// TriangleRight indicates a Triangles where one angle is at 90 degrees
	TriangleRight
)

// Triangle is a basic triangular form; the "point" of the triangle is pointing to the top. The dimensions are controlled
// by the SpaceComponent.
type Triangle struct {
	TriangleType TriangleType

	BorderWidth float32
	BorderColor color.Color
}

// Texture always returns nil. Triangle is drawable without a Texture. This implements the Drawable interface.
func (Triangle) Texture() *gl.Texture { return nil }

// Width always returns 0. This implements the Drawable interface.
func (Triangle) Width() float32 { return 0 }

// Height always returns 0. This implements the Drawable interface.
func (Triangle) Height() float32 { return 0 }

// View always returns 0, 0, 1, 1. This implements the Drawable interface.
func (Triangle) View() (float32, float32, float32, float32) { return 0, 0, 1, 1 }

// Close does nothing, because there's no Texture on the GPU. This implements the Drawable interface.
func (Triangle) Close() {}

// Rectangle is a basic rectangular form; the dimensions are controlled via the `SpaceComponent`.
type Rectangle struct {
	BorderWidth float32
	BorderColor color.Color
}

// Texture always returns nil. Rectangle is drawable without a Texture. This implements the Drawable interface.
func (Rectangle) Texture() *gl.Texture { return nil }

// Width always returns 0. This implements the Drawable interface.
func (Rectangle) Width() float32 { return 0 }

// Height always returns 0. This implements the Drawable interface.
func (Rectangle) Height() float32 { return 0 }

// View always returns 0, 0, 1, 1. This implements the Drawable interface.
func (Rectangle) View() (float32, float32, float32, float32) { return 0, 0, 1, 1 }

// Close does nothing, because there's no Texture on the GPU. This implements the Drawable interface.
func (Rectangle) Close() {}

// Circle is a basic circular form; the dimensions / radius are controlled via the `SpaceComponent`.
// This was made possible by the shared knowledge of Olivier Gagnon (@hydroflame).
type Circle struct {
	BorderWidth float32
	BorderColor color.Color
}

// Texture always returns nil. Circle is drawable without a Texture. This implements the Drawable interface.
func (Circle) Texture() *gl.Texture { return nil }

// Width always returns 0. This implements the Drawable interface.
func (Circle) Width() float32 { return 0 }

// Height always returns 0. This implements the Drawable interface.
func (Circle) Height() float32 { return 0 }

// View always returns 0, 0, 1, 1. This implements the Drawable interface.
func (Circle) View() (float32, float32, float32, float32) { return 0, 0, 1, 1 }

// Close does nothing, because there's no Texture on the GPU. This implements the Drawable interface.
func (Circle) Close() {}

// ComplexTriangles is a complex form, made out of triangles.
type ComplexTriangles struct {
	// Points are the points the form is made of. They should be defined on a scale from 0 to 1, where (0, 0) starts
	// at the top-left of the area (as defined by the `SpaceComponent`.
	// You should use a multitude of 3 points, because each triangle is defined by defining 3 points.
	Points []engo.Point

	// BorderWidth indicates the width of the border, around EACH of the Triangles it is made out of
	BorderWidth float32
	// BorderColor indicates the color of the border, around EACH of the Triangles it is made out of
	BorderColor color.Color
}

// Texture always returns nil. ComplexTriangles is drawable without a Texture. This implements the Drawable interface.
func (ComplexTriangles) Texture() *gl.Texture { return nil }

// Width always returns 0. This implements the Drawable interface.
func (ComplexTriangles) Width() float32 { return 0 }

// Height always returns 0. This implements the Drawable interface.
func (ComplexTriangles) Height() float32 { return 0 }

// View always returns 0, 0, 1, 1. This implements the Drawable interface.
func (ComplexTriangles) View() (float32, float32, float32, float32) { return 0, 0, 1, 1 }

// Close does nothing, because there's no Texture on the GPU. This implements the Drawable interface.
func (ComplexTriangles) Close() {}
