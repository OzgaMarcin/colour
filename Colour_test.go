package colour

import (
	"testing"
)

func TestColour_String(t *testing.T) {
	tests := []struct {
		name   string
		colour Colour
		want   string
	}{
		{
			name:   "White",
			colour: White,
			want:   "FFFFFFFF",
		},
		{
			name:   "Blue",
			colour: Blue,
			want:   "FF0000FF",
		},
		{
			name:   "Red",
			colour: Red,
			want:   "FFFF0000",
		},
		{
			name:   "Green",
			colour: Lime,
			want:   "FF00FF00",
		},
		{
			name:   "Black",
			colour: Black,
			want:   "FF000000",
		},
		{
			name:   "Magenta",
			colour: Magenta,
			want:   "FFFF00FF",
		},
		{
			name:   "Yellow",
			colour: Yellow,
			want:   "FFFFFF00",
		},
		{
			name:   "Aqua",
			colour: Aqua,
			want:   "FF00FFFF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.colour.String(); got != tt.want {
				t.Errorf("Colour.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColour_ToString(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		colour Colour
		args   args
		want   string
	}{
		{
			name:   "ARGB",
			colour: Gray,
			args: args{
				format: FormatARGB,
			},
			want: "FF808080",
		},
		{
			name:   "RGB",
			colour: Gray,
			args: args{
				format: FormatRGB,
			},
			want: "808080",
		},
		{
			name:   "WEB",
			colour: Gray,
			args: args{
				format: FormatWEB,
			},
			want: "#808080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.colour.ToString(tt.args.format); got != tt.want {
				t.Errorf("Colour.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
