package shape

//struct
type Rectangle struct {
	Width  float32
	Length float32
}

//method menghitung luas
func (rectangle *Rectangle) Area() float32 {
	return rectangle.Width * rectangle.Length
}

//method menghitung keliling
func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Width + rectangle.Length)
}
