package utilities

type Point struct {
	x float64
	y float64
}

type Rect struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

// PDF 1.6 Reference: Transformation matrix
type Matrix struct {
	a float64
	b float64
	c float64
	d float64
	e float64
	f float64
}

func GetIdentityMatrix() Matrix {
	return Matrix{
		a: 1,
		b: 0,
		c: 0,
		d: 1,
		e: 0,
		f: 0,
	}
}

func MultiplyMatrix(m1 Matrix, m0 Matrix) Matrix {
	//returns multiplication of two matrices
	return Matrix{
		a: m0.a*m1.a + m0.c*m1.b,
		b: m0.b*m1.a + m0.d*m1.b,
		c: m0.a*m1.c + m0.c*m1.d,
		d: m0.b*m1.c + m0.d*m1.d,
		e: m0.a*m1.e + m0.c*m1.f + m0.e,
		f: m0.b*m1.e + m0.d*m1.f + m0.f,
	}
}

func TranslateMatrix(m Matrix, v Point) Matrix {
	// Translates a matrix by (x, y)
	return Matrix{
		a: m.a,
		b: m.b,
		c: m.c,
		d: m.d,
		e: v.x*m.a + v.y*m.c + m.e,
		f: v.x*m.b + v.y*m.d + m.f,
	}
}

func ApplyMatrixToPoint(m Matrix, v Point) Point {
	// Applies a matrix to a point
	return Point{
		x: m.a*v.x + m.c*v.y + m.e,
		y: m.b*v.x + m.d*v.y + m.f,
	}
}

func ApplyMatrixNorm(m Matrix, v Point) Point {
	// Equivalent to ApplyMatrixToPoint(M, (p,q)) - ApplyMatrixToPoint(M, (0,0))
	return Point{
		x: m.a*v.x + m.c*v.y,
		y: m.b*v.x + m.d*v.y,
	}
}
