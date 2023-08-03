package colour

import "fmt"

func (colour Colour) String() string {
	return fmt.Sprintf("%08X", uint32(colour))
}
