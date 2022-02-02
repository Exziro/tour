package word

import (
	"strings"
	"unicode"
)

//ToUpper 变化为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//ToLower变换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

//下划线单词转大写驼峰单词
func UnderscoreToUpperCamelcase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//下划线转小写驼峰单词
func UnderscoreToLowerCamelcase(s string) string {
	s = UnderscoreToUpperCamelcase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰单词转下划线单词
func CamlecaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)

}
