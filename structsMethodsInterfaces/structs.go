package structs

func Perimeter(rectangle Rectangle) float64 {

	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}
