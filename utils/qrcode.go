package utils

import (
	"encoding/base64"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

func GenB64Qrcode(content string) (string, error) {
	var png []byte
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	var b64Qrcode strings.Builder
	b64Qrcode.WriteString("data:image/png;base64,")
	b64Qrcode.WriteString(base64.StdEncoding.EncodeToString(png))
	return b64Qrcode.String(), nil
}

func GenQrcodeImg(content string, fileName string) error {
	// customize background color and foreground color:
	// qrcode.WriteColorFile(content, qrcode.Medium, 256, color.Black, color.White, fileName)
	return qrcode.WriteFile(content, qrcode.Medium, 256, fileName)
}
