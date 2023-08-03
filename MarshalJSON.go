package colour

import "encoding/json"

func (colour Colour) MarshalJSON() ([]byte, error) {
	return json.Marshal(colour.String())
}
