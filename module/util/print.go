package util

import "encoding/json"

// 高亮颜色map
var colorMap = map[string]string{
	"green":   "\033[97;42m",
	"white":   "\033[90;47m",
	"yellow":  "\033[90;43m",
	"red":     "\033[97;41m",
	"blue":    "\033[97;44m",
	"magenta": "\033[97;45m",
	"cyan":    "\033[97;46m",
	"reset":   "\033[0m",
}

// HighlightString 高亮字符串
func HighlightString(color string, str string) string {
	// 判断是否存在颜色，不存在返回绿色
	if _, ok := colorMap[color]; !ok {
		return colorMap["green"] + str + colorMap["reset"]
	}
	return colorMap[color] + str + colorMap["reset"]
}

func StructToString(s interface{}) string {
	v, _ := json.Marshal(s)
	return string(v)
}
