package utils

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

var driverString = base64Captcha.DriverString{
	Height:          40,
	Width:           100,
	NoiseCount:      0,
	ShowLineOptions: 2 | 4,
	Length:          4,
	Source:          "23456789abcdefghijkmnpqrstuvwxyz", // remove indistinguishable chars 0 o 1 l
	BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
	Fonts:           []string{"wqy-microhei.ttc"},
}
var driver = driverString.ConvertFonts()

// question is the image with chars and noise, the image is encoded into base64 string.
// answer is the chars in the image.
// in order to verify the answer, you need to record the mapping between id and answer.
func NewBase64Captcha() (id, answer, b64Captcha string, err error) {
	id, _, answer = driver.GenerateIdQuestionAnswer() // content equals to answer here
	item, err := driver.DrawCaptcha(answer)
	if err != nil {
		return "", "", "", err
	}
	b64Captcha = item.EncodeB64string()
	return
}

func GenerateBaseCaptchaFromAnswer(answer string) (b64Captcha string, err error) {
	item, err := driver.DrawCaptcha(answer)
	if err != nil {
		return "", err
	}
	b64Captcha = item.EncodeB64string()
	return
}
