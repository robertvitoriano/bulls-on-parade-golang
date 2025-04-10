package utils

const (
	ScreenWidth  = 320
	ScreenHeight = 256
	FPS          = 10
	SCALE        = 2
)

type Position struct {
	X float64
	Y float64
}
type Size struct {
	Width  float64
	Height float64
}
type Velocity struct {
	X float64
	Y float64
}
type CollisionSide string

const (
	CollisionNone   CollisionSide = "NONE"
	CollisionRight  CollisionSide = "RIGHT"
	CollisionLeft   CollisionSide = "LEFT"
	CollisionTop    CollisionSide = "TOP"
	CollisionBottom CollisionSide = "BOTTOM"
)
