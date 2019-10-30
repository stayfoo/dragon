/*
@Time : 2019-10-26 12:53
@Author : mengyueping
@File : parser
@Software: GoLand
*/
package parser

import (
	"dragon/lexer"
	"fmt"
	"log"
)

/**
一个语法解析器。
能够解析简单的表达式、变量声明和初始化语句、赋值语句。
它支持的语法规则：

additionSubtractionExpression  //加法、减法的表达式的规则
    : multiplyDivideExpression //乘法、除法的表达式规则，优先级高于加法、减法。
    | multiplyDivideExpression Add additionSubtractionExpression //表达式 + 表达式
    | multiplyDivideExpression Sub additionSubtractionExpression //表达式 - 表达式
    ;
multiplyDivideExpression  //乘法、除法的表达式的规则
    : primary_expression  //一元表达式
    | primary_expression Mul multiplyDivideExpression //表达式 * 表达式
    | primary_expression Div multiplyDivideExpression //表达式 / 表达式
    ;
primary_expression //一元表达式的规则
    : IntLiteral   //int的字面量
    ;


program -> intDeclare | expressionStatement | assignmentStatement
	intDeclare -> 'int' Id ( = add) ';'
	expressionStatement -> add ';'
		add -> multiply ( (+ | -) multiply)*
			multiply -> primary ( (* | /) primary)*
				primary -> IntLiteral | Id | (add)
	assignmentStatement -> Id = add ';'
		add -> multiply ( (+ | -) multiply)*
			multiply -> primary ( (* | /) primary)*
				primary -> IntLiteral | Id | (add)
*/

type Parser struct {
	Lexer *lexer.Lexer
}

//解析代码
func (p *Parser)Parse(code string, verbose bool) *ASTNode {
	if p.Lexer == nil {
		l := &lexer.Lexer{}
		l.Tokenize(code)
		if verbose {
			l.Reader.Dump()
		}
		p.Lexer = l
	}

	n := p.program()
	if verbose {
		fmt.Println("AST:<< ")
		n.DumpAST("")
		fmt.Println(">>")
	}

	return n
}

//AST的根节点，语法解析的入口
func (p *Parser)program() *ASTNode {
	r := p.Lexer.Reader
	node := &ASTNode{Type: Program, Text: "pwc"}
	for r.Peek() != nil {
		child := p.intDeclare()

		if child == nil {
			child = p.expressionStatement()
		}

		if child == nil {
			child = p.assignmentStatement()
		}

		if child != nil {
			node.AddChild(child)
		} else {
			//语法错误：unknown statement
			log.Fatalln("unknown statement expression.")
		}
	}
	return node
}

//表达式语句，表达式后面跟个分号
func (p *Parser)expressionStatement() *ASTNode {
	r := p.Lexer.Reader
	pos := r.Position
	node := p.add() //可能会对 r 做 Read() 操作，导致 Position 改变
	if node != nil {
		t := r.Peek()
		if t != nil && t.Type == lexer.SemiColon { // ;
			r.Read()
		} else {
			node = nil
			r.Position = pos //回溯，以防万一，可以不做，此分支前面没有执行 Read() 操作，没有改变 Position
		}
	}
	return node //直接返回子节点
}

//赋值语句，例如：
// age = 10;
func (p *Parser)assignmentStatement() *ASTNode {
	r := p.Lexer.Reader
	var n *ASTNode
	t := r.Peek()
	if t != nil && t.Type == lexer.Identifier {
		t = r.Read() //标识符
		n = &ASTNode{Type: AssignmentStmt, Text: t.Text}
		t = r.Peek()
		if t != nil && t.Type == lexer.Assignment {
			r.Read() //赋值符号
			child := p.add()
			if child == nil {
				//语法错误：等号右边没有合法的表达式
				log.Fatalln("invalid assignmentStatement expression, expecting the right is correct.")
			} else {
				n.AddChild(child) //添加子节点
				t = r.Peek()
				if t != nil && t.Type == lexer.SemiColon { // ;
					r.Read()
				} else {
					//语法错误：没有结束符;号
					log.Fatalln("invalid assignmentStatement expression, expecting the end of SemiColon.")
				}
			}
		} else {
			r.Unread() //回溯，重置之前消化掉的标识符
			n = nil
		}
	}
	return n
}

//整型变量声明，例如：
// int a;
// int b = 2 * 3;
func (p *Parser)intDeclare() *ASTNode {
	var n *ASTNode
	r := p.Lexer.Reader
	t := r.Peek()
	if t != nil && t.Type == lexer.Int { // int 类型
		t = r.Read()
		if r.Peek().Type == lexer.Identifier { // 标识符
			t = r.Read()
			n = &ASTNode{Type: IntDeclaration, Text: t.Text}
			t = r.Peek()
			if t != nil && t.Type == lexer.Assignment { // =
				r.Read()
				child := p.add()
				if child == nil {
					//语法错误：变量初始化错误
				} else {
					n.AddChild(child)
				}
			}
		} else {
			//语法错误：变量名字异常
			log.Fatalln("invalid intDeclare expression, expecting the right name.")
		}

		if n != nil {
			t = r.Peek()
			if t != nil && t.Type == lexer.SemiColon {
				r.Read()
			} else {
				//语法错误：位置的代码片段
				log.Fatalln("invalid intDeclare expression, expecting the end of SemiColon.")
			}
		}
	}
	return n
}

//加法运算
func (p *Parser)add() *ASTNode {
	r := p.Lexer.Reader
	child1 := p.multiply() //计算第一个子节点
	n := child1           //第二个子节点，如果没有，就返回第一个
	if child1 != nil {
		for {
			t := r.Peek()
			if t != nil && (t.Type == lexer.Plus || t.Type == lexer.Minus) { // + -
				t = r.Read()
				child2 := p.multiply() //递归解析第二个节点
				if child2 != nil {
					n = &ASTNode{Type: Additive, Text: t.Text}
					n.AddChild(child1) //注意：新节点在顶层，保证正确的结合性
					n.AddChild(child2)
					child1 = n
				} else {
					log.Fatalln("invalid additive expression, expecting the right part.")
				}
			} else {
				break
			}
		}
	}
	return n
}

//乘法
func (p *Parser)multiply() *ASTNode {
	r := p.Lexer.Reader
	child1 := p.primary()
	n := child1

	for {
		t := r.Peek()
		if t != nil && (t.Type == lexer.Star || t.Type == lexer.Slash) { // * /
			t = r.Read()
			child2 := p.primary()
			if child2 != nil {
				n = &ASTNode{Type: Multiplicative, Text: t.Text}
				n.AddChild(child1)
				n.AddChild(child2)
				child1 = n
			} else {
				//语法错误：乘法|除法表达式右边部分
				log.Fatalln("invalid multiply expression, expecting the right part.")
			}
		} else {
			break
		}
	}

	return n
}

//基础表达式
func (p *Parser)primary() *ASTNode {
	r := p.Lexer.Reader
	var n *ASTNode
	t := r.Peek()

	if t != nil {
		if t.Type == lexer.IntLiteral {
			t = r.Read()
			n = &ASTNode{Type: IntLiteral, Text: t.Text}
		} else if t.Type == lexer.Identifier {
			t = r.Read()
			n = &ASTNode{Type: Identifier, Text: t.Text}
		} else if t.Type == lexer.LeftParen {
			r.Read()
			n = p.add() //匹配加减乘除算术表达式，返回下一个节点
			if n != nil {
				t = r.Peek()
				if t != nil && t.Type == lexer.RightParen {
					r.Read()
				} else {
					//语法错误：缺少右括号
					log.Fatalln("invalid primary expression, expecting the right part.")
				}
			} else {
				//语法错误：括号中缺少算术表达式
				log.Fatalln("invalid primary expression.")
			}
		}
	}
	return n //做了AST的简化，不用构造一个 primary 节点，直接返回子节点。它只有一个子节点
}
