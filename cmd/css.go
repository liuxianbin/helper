package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"helper/internal/css"
	"strings"
)

const (
	AUTO = iota // 自动转换成HEX或RGB
	MODE_HEX_TO_RGB
	MODE_RGB_TO_HEX
)

var desc = strings.Join([]string{
	"CSS颜色值转换,模式如下:",
	"0: 自动转换成HEX或RGB",
	"1: HEX转换成RGB",
	"2: RGB转换成HEX",
}, "\n")

var cssCmd = &cobra.Command{
	Use:   "css",
	Short: "CSS颜色值转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		if content == "" {
			return
		}
		var err error
		switch mode {
		case AUTO:
			if strings.ToLower(color[:3]) == "rgb" {
				content, err = css.RGB_TO_HEX(color)
			} else {
				content, err = css.HEX_TO_RGB(color)
			}
		case MODE_HEX_TO_RGB:
			content, err = css.HEX_TO_RGB(color)
		case MODE_RGB_TO_HEX:
			content, err = css.RGB_TO_HEX(color)
		}
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}
		fmt.Println(content)
	},
}

// 颜色值
var color string

// 颜色值转换的模式
var mode int8

func init() {
	cssCmd.Flags().StringVarP(&color, "color", "c", "", "颜色值")
	cssCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "颜色值转换的模式")
}
