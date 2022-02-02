package cmd

import (
	"log"
	"strings"

	"github.com/go_tour/tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1 //
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下:",
	"1:全部单词转为大写",
	"2:全部单词转为小写",
	"3:下画线单词转为大写驼峰单词",
	"4:下画线单词转为小写驼峰单词",
	"5:驼峰单词转为下画线单词"}, "\n")
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式变换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelcase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelcase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamlecaseToUnderscore(str)
		default:
			log.Fatal("暂不支持该转换模式，请执行help word 查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}
var str string
var mode int8

func init() {
	/*Varp函数
	第一个参数为需要绑定的变量，
	第二个参数为接收该参数的完整的命令标志，
	第三个参数为对应的短标识，
	第四个参数为默认值，
	第五个参数为使用说明。*/
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入转换的模式")
}
