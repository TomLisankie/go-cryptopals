package main

import "testing"

func TestHexStringToBase64String(t *testing.T) {
	testHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	testBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got, err := HexStringToBase64String(testHex)
	if err != nil {
		t.Errorf("There was an error")
	}
	if got != testBase64 {
		t.Errorf("Did not produce the correct base64 string")
	}
}
