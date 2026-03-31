package utils

import "github.com/skip2/go-qrcode"

func GenerateQR(data string, path string) error {
	return qrcode.WriteFile(data, qrcode.Medium, 256, path)
}
