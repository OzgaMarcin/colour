package colour

func FromARGB(a byte, r byte, g byte, b byte) Colour {
	return Colour(uint32(a)<<24 | uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}
