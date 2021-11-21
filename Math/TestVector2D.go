package Math

import (
	"testing"
)

func TestNew(t *testing.T) {
	var x = float64(1.1)
	var y = float64(2.2)

	v := NewVector2D(x, y)
	if v.X != x || v.Y != y {
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	v := ZeroVector2d()
	if v.X != 0 || v.Y != 0 {
		t.Fail()
	}
}

func TestUnit(t *testing.T) {
	v := UnitVector2D()
	if v.X != 1 || v.Y != 1 {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	var x = float64(1.1)
	var y = float64(2.2)

	v1 := Vector2D{x, y}
	v2 := v1.Copy()
	if v2.X != x || v2.Y != y {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	var x1 = float64(1)
	var y1 = float64(2)
	var x2 = float64(4)
	var y2 = float64(8)

	v1 := Vector2D{x1, y1}
	v2 := Vector2D{x2, y2}
	result := v1.Add(v2)
	if result.X != 5 || result.Y != 10 {
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	var x1 = float64(1)
	var y1 = float64(2)
	var x2 = float64(4)
	var y2 = float64(8)

	v1 := Vector2D{x1, y1}
	v2 := Vector2D{x2, y2}
	result := v1.Subtract(v2)
	if result.X != -3 || result.Y != -6 {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	var x1 = float64(1)
	var y1 = float64(2)
	var x2 = float64(4)
	var y2 = float64(8)

	v1 := Vector2D{x1, y1}
	v2 := Vector2D{x2, y2}
	result := v1.Multiply(v2)
	if result.X != 4 || result.Y != 16 {
		t.Fail()
	}
}

func TestDivide(t *testing.T) {
	var x1 = float64(1)
	var y1 = float64(2)
	var x2 = float64(4)
	var y2 = float64(10)

	v1 := Vector2D{x1, y1}
	v2 := Vector2D{x2, y2}
	result := v1.Divide(v2)
	if result.X != 0.25 || result.Y != 0.2 {
		t.Fail()
	}
}

func TestMultiplyScalar(t *testing.T) {
	var x = float64(1)
	var y = float64(2)

	v := Vector2D{x, y}
	result := v.MultiplyScalar(1.1)
	if result.X != 1.1 || result.Y != 2.2 {
		t.Fail()
	}
}

func TestDivideScalar(t *testing.T) {
	var x = float64(1)
	var y = float64(2)

	v := Vector2D{x, y}
	result := v.DivideScalar(2)
	if result.X != 0.5 || result.Y != 1 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var x = float64(1)
	var y = float64(2)

	v := Vector2D{x, y}
	result := v.String()
	if result != "1:2" {
		t.Fail()
	}
}
