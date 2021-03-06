package girya

import (
	"strconv"
	"testing"
)

func TestWeight_String(t *testing.T) {
	tests := []struct {
		weight Weight
		want   string
	}{
		{2020, "2.02 t"},
		{2004, "2 t"},
		{-2020, "-2.02 t"},
		{2000, "2 t"},
		{2228000, "2228 t"},
		{2100, "2.1 t"},
		{22000, "22 t"},
		{22030, "22.03 t"},
		{22333, "22.33 t"},
		{22300, "22.3 t"},
		{2000, "2 t"},
		{2022, "2.02 t"},
		{1990, "1.99 t"},
		{8, "8 kg"},
		{80190, "80 t"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.weight)), func(t *testing.T) {
			if got := tt.weight.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
