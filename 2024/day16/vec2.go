package main

type Vec2 struct {
	X int
	Y int
}

func (i Vec2) LessThan(j Vec2) bool {
	return (i.X*i.X + i.Y*i.Y) < (j.X*j.X + j.Y*j.Y)
}

func (i Vec2) EqualTo(j Vec2) bool {
	return i.X == j.X && i.Y == j.Y
}

func vec2(x, y int) Vec2 {
	return Vec2{X: x, Y: y}
}
