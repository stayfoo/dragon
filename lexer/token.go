/*
@Time : 2019-10-26 12:50
@Author : mengyueping
@File : token
@Software: GoLand
*/
package lexer

type Token struct {
	Type TokenType
	Text string
}

var TokenMap = map[string]string{
	"Id":         `[a-zA-Z_] ([a-zA-Z_] | [0-9])*`,
	"IntLiteral": `[0-9]+`,
	"GT":         `'>'`,
	"GE":         `'>='`,
	"Int":        `'int'`,
	"Assignment": `'='`,
}
