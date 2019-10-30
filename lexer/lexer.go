/*
@Time : 2019-10-26 12:53
@Author : mengyueping
@File : lexer
@Software: GoLand
*/
package lexer

//一个手写的词法分析器。
//能够为后面的简单计算器、简单脚本语言产生Token。
type Lexer struct {
	Reader *TokenReader
}

//Deterministic finite state machine, 确定的有限状态机
//有限状态机的各种状态
type FiniteState int

const (
	FiniteStateInitial FiniteState = iota

	FiniteStateIf
	FiniteStateIdIf1
	FiniteStateIdIf2

	FiniteStateElse
	FiniteStateIdElse1
	FiniteStateIdElse2
	FiniteStateIdElse3
	FiniteStateIdElse4

	FiniteStateInt
	FiniteStateIdInt1
	FiniteStateIdInt2
	FiniteStateIdInt3

	FiniteStateId
	FiniteStateGT //>
	FiniteStateGE //>=
	FiniteStateEQ //==

	FiniteStateAssignment //=

	FiniteStatePlus
	FiniteStateMinus
	FiniteStateStar
	FiniteStateSlash

	FiniteStateSemiColon  // ;
	FiniteStateLeftParen  // (
	FiniteStateRightParen // )

	FiniteStateSpace

	FiniteStateIntLiteral
)

//有限状态机进入初始状态。
//在初始状态不做停留，马上进入其他状态。
//开始解析的时候，进入初始状态；某个Token解析完毕，也进入初始状态，在这里把Token记下来，然后建立一个新的Token。
func (l *Lexer)initToken(r rune) (Token, FiniteState) {
	t := Token{}
	s := string(r)
	state := FiniteStateInitial
	if IsAlpha(r) { //是否是字母
		if s == "i" {
			state = FiniteStateIdInt1
		} else {
			state = FiniteStateId
		}
		t.Type = Identifier
	} else if IsDigit(r) { //是否是数字
		state = FiniteStateIntLiteral
		t.Type = IntLiteral
	} else if r == '>' {
		state = FiniteStateGT
		t.Type = GT
	} else if s == "+" {
		state = FiniteStatePlus
		t.Type = Plus
	} else if s == "-" {
		state = FiniteStateMinus
		t.Type = Minus
	} else if s == "*" {
		state = FiniteStateStar
		t.Type = Star
	} else if s == "/" {
		state = FiniteStateSlash
		t.Type = Slash
	} else if s == ";" {
		state = FiniteStateSemiColon
		t.Type = SemiColon
	} else if s == "(" {
		state = FiniteStateLeftParen
		t.Type = LeftParen
	} else if s == ")" {
		state = FiniteStateRightParen
		t.Type = RightParen
	} else if s == "=" {
		state = FiniteStateAssignment
		t.Type = Assignment
	} else if IsSpace(r) {
		state = FiniteStateSpace
		t.Type = Space
		s = ""
	} else {
		state = FiniteStateInitial //未知状态，跳过
		s = ""
	}
	t.Text = s

	return t, state
}

//解析字符串，形成Token
//这是一个有限状态自动机，在不同的状态中的迁移。
func (l *Lexer)Tokenize(code string) TokenReader {
	runeList := []rune(code)
	var tokenList []Token

	t := Token{}
	state := FiniteStateInitial

	for i := 0; i < len(runeList); i++ {
		r := runeList[i]
		s := string(r)
		switch state {
		case FiniteStateInitial: //重新确定后续状态
			t, state = l.initToken(r)
		case FiniteStateId:
			if IsAlpha(r) || IsDigit(r) {
				t.Text += s //token保持标识符type
			} else { //退出标识符状态，并保存 Token
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			}
		case FiniteStateGT:
			if s == "=" { //状态转换
				state = FiniteStateGE
				t.Text += s
				t.Type = GE
			} else { //退出 GT 状态，并保存 Token
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			}
		case FiniteStateAssignment:
			if IsSpace(r) {
				t.Type = Assignment
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			} else if s == "=" { //EQ // ==
				state = FiniteStateEQ
				t.Text += s
				t.Type = EQ
			} else {
				//语法错误
			}
		case FiniteStateEQ: //==
			if IsSpace(r) {
				t.Type = EQ
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			} else {
				//语法错误
			}
		case FiniteStateGE: //退出当前状态，并保存 Token
			fallthrough
		case FiniteStatePlus:
			fallthrough
		case FiniteStateMinus:
			fallthrough
		case FiniteStateStar:
			fallthrough
		case FiniteStateSlash:
			fallthrough
		case FiniteStateSemiColon:
			fallthrough
		case FiniteStateLeftParen:
			fallthrough
		case FiniteStateRightParen:
			tokenList = append(tokenList, t)
			t, state = l.initToken(r)
		case FiniteStateSpace:
			t, state = l.initToken(r)
		case FiniteStateIntLiteral:
			if IsDigit(r) { //继续保持在数字字面量状态
				t.Text += s
			} else if IsSpace(r) || s == ";" || s == "," {
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			} else { //退出 IntLiteral 状态，并保存 Token
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			}
		case FiniteStateIdInt1:
			if s == "n" {
				state = FiniteStateIdInt2
				t.Text += s //token保持标识符type
			} else if IsAlpha(r) || IsDigit(r) {
				t.Text += s           //token保持标识符type
				state = FiniteStateId //切回到Id状态
			} else { //退出标识符状态，并保存 Token
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			}
		case FiniteStateIdInt2:
			if s == "t" {
				state = FiniteStateIdInt3
				t.Text += s
			} else if IsAlpha(r) || IsDigit(r) {
				state = FiniteStateId //切回到Id状态
				t.Text += s           //token保持标识符type
			} else { //退出标识符状态，并保存 Token
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			}
		case FiniteStateIdInt3:
			if IsSpace(r) {
				t.Type = Int
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			} else if IsAlpha(r) || IsDigit(r) {
				state = FiniteStateId //切回到Id状态
				t.Text += s           //token保持标识符type
			} else { //退出标识符状态，并保存 Token
				tokenList = append(tokenList, t)
				t, state = l.initToken(r)
			}
		}

		if i == len(runeList)-1 && len(t.Text) > 0 { //保存最后一个 Token
			tokenList = append(tokenList, t)
		}
	}

	r := TokenReader{TokenList: tokenList}
	l.Reader = &r
	return r
}


