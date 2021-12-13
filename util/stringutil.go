package util

import "strings"

const (
	SignBlank                     = ""                               // 空
	SignCarriageReturn            = "\r"                             // 回车
	SignNewLine                   = "\n"                             // 换行
	SingNewLineWithCarriageReturn = SignCarriageReturn + SignNewLine // 回车换行
)

var allSignArr = getAllSignArr()

// TrimEscapeSign 去除特殊符号
func TrimEscapeSign(str string) string {
	retStr := TrimSpace(str)
	if IsBlank(retStr) {
		return str
	}
	retStr = TrimSignWithBlank(str, allSignArr...)

	return retStr
}

// TrimSignWithBlank 去除指定的特殊符号
func TrimSignWithBlank(str string, signs ...string) string {
	if len(signs) == 0 {
		return str
	}

	for _, sign := range signs {
		if IsBlank(str) {
			return str
		}
		str = strings.ReplaceAll(str, sign, SignBlank)
	}

	return str
}

// TrimSpace 去除空格
func TrimSpace(str string) string {
	if IsBlank(str) {
		return str
	}
	return strings.TrimSpace(str)
}

// IsBlank 是否为空(空格也算空)
func IsBlank(str string) bool {
	if IsEmpty(str) || IsEmpty(strings.TrimSpace(str)) {
		return true
	}
	return false
}

// IsEmpty 是否为空(空格不算空)
func IsEmpty(str string) bool {
	return len(str) == 0
}

// 获取所有特殊字符
func getAllSignArr() []string {
	return []string{SignBlank, SignCarriageReturn, SignNewLine, SingNewLineWithCarriageReturn}
}
