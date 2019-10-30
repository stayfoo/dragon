/*
@Time : 2019-10-26 12:49
@Author : mengyueping
@File : token_type
@Software: GoLand
*/
package lexer

type TokenType int

const (
	Initial TokenType = iota //初始状态
	Space                    //空格

	Plus  // +
	Minus // -
	Star  // *
	Slash // /

	GE // >=
	GT // >
	EQ // ==
	LE // <=
	LT // <

	SemiColon  // ;
	EndOfLine  // 一行结束符号：换行
	LeftParen  // (
	RightParen // )

	Assignment // =

	If
	Else

	Int

	Identifier //标识符

	IntLiteral    //整型字面量
	StringLiteral //字符串字面量
)

var TokenTypeName = map[TokenType]string{
	Initial: "Initial",
	Space:   "Space",

	Plus:  "Plus",
	Minus: "Minus",
	Star:  "Star",
	Slash: "Slash",

	GT: "GT",
	GE: "GE",
	EQ: "EQ", // ==
	LE: "LE", // <=
	LT: "LT", // <

	SemiColon:  "SemiColon",  // ;
	EndOfLine:  "EndOfLine",  //一行结束符号：换行
	LeftParen:  "LeftParen",  // (
	RightParen: "RightParen", // )

	Assignment: "Assignment", // =

	If:   "If",
	Else: "Else",

	Int: "Int",

	Identifier: "Identifier",

	IntLiteral: "IntLiteral",

	StringLiteral: "StringLiteral",
}
