// 安全加固（验证码与防刷）登录验证码
package captcha

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/dchest/captcha"
	"github.com/golang/freetype/truetype"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	captcha.New(6, captcha.NumDigits) // 6位数字验证码
}

// 获取验证码图片
func GenerateCaptcha() (string, []byte) {
	id := captcha.NewID()
	return id, captcha.NewImage(id, 200, 80, &captcha.Options{
		Digits: 6,
		Colors: &captcha.ColorOptions{
			BgColor: color.White,
			FgColor: color.Black,
			Font:    truetype.NewFace(nil, &truetype.Options{}),
		},
	})
}

// 验证验证码
func VerifyCaptcha(id, solution string) bool {
	return captcha.Verify(id, solution)
}
