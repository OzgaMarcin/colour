package colour

import "fmt"

const (
	FormatARGB = "%08X"
	FormatRGB  = "%06X"
	FormatWEB  = "#%06X"
)

func (colour Colour) ToString(format string) string {
	switch format {
	case FormatARGB:
		return fmt.Sprintf(format, uint32(colour))
	case FormatRGB:
		return fmt.Sprintf(format, uint32(colour)&0x00FFFFFF)
	case FormatWEB:
		return fmt.Sprintf(format, uint32(colour)&0x00FFFFFF)
	default:
		return fmt.Sprintf(format, uint32(colour))
	}
}
