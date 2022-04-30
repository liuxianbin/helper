package css

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rgb struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (r rgb) String() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", r.Red, r.Green, r.Blue)
}

func (r rgb) ToHEX() string {
	return strings.Join([]string{
		"#",
		fmt.Sprintf("%x", r.Red),
		fmt.Sprintf("%x", r.Green),
		fmt.Sprintf("%x", r.Blue),
	}, "")
}

var InvalidHEX = errors.New("无效的HEX值")
var InvalidRGB = errors.New("无效的RGB值")

func HEX_TO_RGB(s string) (string, error) {
	if s[0] == '#' {
		s = s[1:]
	}
	if len(s) > 6 {
		return "", InvalidHEX
	}
	s += strings.Repeat("0", len(s)-6)
	n, err := strconv.ParseInt(s, 16, 0)
	if err != nil {
		return "", InvalidHEX
	}
	return rgb{
		Red:   uint8(n >> 16),
		Green: uint8(n & 0x00ff00 >> 8),
		Blue:  uint8(n & 0x0000ff),
	}.String(), nil
}

func RGB_TO_HEX(s string) (string, error) {
	e, _ := regexp.Compile(`^rgb\((\d{1,3}),(\d{1,3}),(\d{1,3})\)$`)
	result := e.FindStringSubmatch(s)
	if len(result) == 4 {
		red, _ := strconv.Atoi(result[1])
		green, _ := strconv.Atoi(result[2])
		blue, _ := strconv.Atoi(result[3])
		return rgb{
			Red:   uint8(red),
			Green: uint8(green),
			Blue:  uint8(blue),
		}.ToHEX(), nil
	}
	return "", InvalidRGB
}
