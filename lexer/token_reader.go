/*
@Time : 2019-10-26 12:51
@Author : mengyueping
@File : token_reader
@Software: GoLand
*/
package lexer

import "fmt"

//一个Token流。把Token列表进行了封装。Parser可以从中获取Token。
type TokenReader struct {
	TokenList []Token
	//流当前的读取位置
	Position int
}

//返回Token流中下一个Token，但不从流中取出。如果流已经为空，返回nil。
func (r *TokenReader) Peek() *Token {
	if r.Position < len(r.TokenList) {
		t := r.TokenList[r.Position]
		return &t
	}
	return nil
}

//返回Token流中下一个Token，并从流中取出。如果流已经为空，返回nil。
func (r *TokenReader) Read() *Token {
	if r.Position < len(r.TokenList) {
		t := r.TokenList[r.Position]
		r.Position++
		return &t
	}
	return nil
}

//Token流回退一步。恢复原来的Token。
func (r *TokenReader) Unread() {
	if r.Position > 0 {
		r.Position--
	}
}

//打印所有Token
func (r *TokenReader) Dump() {
	pos := r.Position

	fmt.Println("token:<< ")
	fmt.Println("text\t|type\t\t|")
	for token := r.Read(); token != nil; token = r.Read() {
		fmt.Println(token.Text + "\t" + TokenTypeName[token.Type])
	}
	fmt.Println(">>")

	r.Position = pos
}


