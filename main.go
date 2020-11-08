package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func HexStringToBase64String(hexString string) (string, error) {
	theBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", errors.New("There was an error decoding the hex string")
	}
	base64String := base64.StdEncoding.EncodeToString(theBytes)
	return base64String, nil
}
