package main

import "testing"

const (
	input  = "Yo Yo Yoshita dudes. Lets say hello-world_2, we like using camelCase more than snake_case."
	output = "Cs Cs Cswimxe hyhiw. Pixw wec lipps-asvph_6, ai pmoi ywmrk geqipGewi qsvi xler wreoi_gewi."
)

func TestEncrypt(t *testing.T) {
	if encrypt(input) != output {
		t.Fatal("error  ", encrypt(input))
	}
}
