package colour

func FromRGB(r byte, g byte, b byte) Colour {
	return FromARGB(0xFF, r, g, b)
}
