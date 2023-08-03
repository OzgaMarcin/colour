package colour

import (
	"strconv"
	"strings"
)

func (colour *Colour) UnmarshalJSON(data []byte) error {
	colourString := strings.Trim(string(data), "\"")
	if colourString == "" {
		*colour = 0
		return nil
	}
	colourUint64, err := strconv.ParseUint(colourString, 16, 64)
	if err != nil {
		return err
	}
	*colour = Colour(colourUint64)
	return nil
}
