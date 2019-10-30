/*
@Time : 2019-10-30 10:28 
@Author : mengyueping
@File : graphic
@Software: GoLand
*/
package lexer

import "unicode"

//判断是否是字母
func IsAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

//判断是否是数字
func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

//判断是否为空白符号: '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0
func IsSpace(r rune) bool {
	return unicode.IsSpace(r)
}