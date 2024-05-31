package utilities

type Rect struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func GetBound(upper_left Point, lower_right Point) Rect {
	return Rect{
		x1: upper_left.x,
		y1: upper_left.y,
		x2: lower_right.x,
		y2: lower_right.y,
	}
}
