package captcha

import (
	"image/png"
	"io"

	"github.com/dchest/captcha"
)

func GenerateCaptcha() (string, io.Reader, error) {
	id := captcha.New() // 无需参数
	solution := captcha.RandomDigits(6) // 生成6位数字
	captcha.Set(id, solution, time.Minute)

	img := captcha.NewImage(id, 200, 80) // 参数正确顺序：id, width, height
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	return id, buf, nil
}

func VerifyCaptcha(id, solution string) bool {
	return captcha.Verify(id, solution) // 直接对比字符串
}