package main

import (
	"encoding/json"
	"testing"
)

func TestIsIMDeliveryContract(t *testing.T) {
	for _, name := range []string{"IM2609", "IM2703"} {
		if !isIMDeliveryContract(name) {
			t.Errorf("expected %q to be an IM delivery contract", name)
		}
	}
	for _, name := range []string{"IM主力合约", "IC2609", "IM261"} {
		if isIMDeliveryContract(name) {
			t.Errorf("expected %q not to be an IM delivery contract", name)
		}
	}
}

func TestParseEastmoneyNumber(t *testing.T) {
	for _, test := range []struct {
		name  string
		value string
		want  float64
		valid bool
	}{
		{name: "json number", value: `729910`, want: 729910, valid: true},
		{name: "json string", value: `"7299.1"`, want: 7299.1, valid: true},
		{name: "unavailable", value: `"-"`, valid: false},
		{name: "null", value: `null`, valid: false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, valid := parseEastmoneyNumber(json.RawMessage(test.value))
			if valid != test.valid || (valid && got != test.want) {
				t.Fatalf("parseEastmoneyNumber(%s) = %v, %v; want %v, %v", test.value, got, valid, test.want, test.valid)
			}
		})
	}
}