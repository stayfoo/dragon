/*
@Time : 2019-10-29 13:29
@Author : mengyueping
@File : calculator
@Software: GoLand
*/
package calculator

import (
	"dragon/lexer"
	"dragon/parser"
	"errors"
	"fmt"
	"log"
	"strconv"
)

/**
实现一个计算器，但计算的结合性是有问题的。
因为它使用了下面的语法规则：

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

program -> add
	multiply | multiply + add
		multiply -> primary | primary * multiply

递归项在右边，会自然的对应右结合。真正需要的是左结合。
*/

type Calculator struct {
	Code   string
	Lexer *lexer.Lexer
	RootNode *parser.ASTNode
	Verbose bool //是否开启日志打印
	Result int //计算结果
	Variables map[string]int //支持定义变量
}

//执行脚本，打印输出AST和求值过程
func (c *Calculator) Evaluate() int {
	c.parse()
	if c.Verbose {
		fmt.Println("AST:<<")
		c.RootNode.DumpAST("")
		fmt.Println(">>")
	}
	if c.Verbose {
		fmt.Println("calculating...")
	}

	result := c.evaluate(c.RootNode, "")
	c.Result = result
	if c.Verbose {
		fmt.Println( "calculate end.")
	}

	return result
}

//解析算式，返回根节点
func (c *Calculator) parse() *parser.ASTNode {
	if c.Lexer == nil {
		l := &lexer.Lexer{}
		l.Tokenize(c.Code)
		if c.Verbose {
			l.Reader.Dump()
		}
		c.Lexer = l
	}
	c.RootNode = c.program()
	return c.RootNode
}

//语法解析：根节点
func (c *Calculator) program() *parser.ASTNode {
	node := &parser.ASTNode{Type: parser.Program, Text: "Calculator"}

	child := c.add()

	if child != nil {
		node.AddChild(child)
	}
	return node
}

//对某个AST节点求值，并打印求值过程
// indent 打印输出时的缩进量，用tab控制
func (c *Calculator) evaluate(n *parser.ASTNode, indent string) int {
	var result int
	if c.Verbose {
		fmt.Println(indent + "calculating:<< " + parser.ASTNodeTypeNameMap[n.Type])
	}

	switch n.Type {
	case parser.Program:
		for _, child := range n.Children {
			result = c.evaluate(child, "\t")
		}
	case parser.Additive:
		child1 := n.Children[0]
		value1 := c.evaluate(child1, "\t")

		child2 := n.Children[1]
		value2 := c.evaluate(child2, "\t")

		if n.Text == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case parser.Multiplicative:
		child1 := n.Children[0]
		value1 := c.evaluate(child1, "\t")

		child2 := n.Children[1]
		value2 := c.evaluate(child2, "\t")

		if n.Text == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case parser.IntLiteral:
		i, err := strconv.Atoi(n.Text)
		if errors.Unwrap(err) != nil {
			log.Fatalln("invalid IntLiteral.")
		} else {
			result = i
		}
	case parser.Identifier:
		varName := n.Text
		if value, ok := c.Variables[varName]; ok {
			result = value
		}else {
			fmt.Println("unknown variable: " + varName)
		}
	case parser.AssignmentStmt:
		varName := n.Text
		if _, ok := c.Variables[varName]; !ok {
			fmt.Println("unknown variable: " + varName)
		}
	case parser.IntDeclaration:
		varName := n.Text
		if len(n.Children)>0 {
			child := n.Children[0]
			result = c.evaluate(child, indent + "\t")
			c.Variables[varName] = result
		}
	default:

	}

	if c.Verbose {
		fmt.Println(indent + "Result: " + fmt.Sprintf("%d", result))
		fmt.Println(indent + "calculate end. >>")
	}

	return result
}

//语法解析：加法表达式
func (c *Calculator) add() *parser.ASTNode {
	r := c.Lexer.Reader
	child1 := c.multiply()
	n := child1

	t := r.Peek()
	if child1 != nil && t != nil {
		if t.Type == lexer.Plus || t.Type == lexer.Minus {
			t = r.Read()
			child2 := c.add()
			if child2 != nil {
				n = &parser.ASTNode{Type: parser.Additive, Text: t.Text}
				n.AddChild(child1)
				n.AddChild(child2)
			} else {
				log.Fatalln("invalid additive expression, expecting the right part.")
			}
		}
	}

	return n
}

//语法解析：乘法表达式
func (c *Calculator) multiply() *parser.ASTNode {
	r := c.Lexer.Reader
	child1 := c.primary()
	n := child1

	t := r.Peek()
	if child1 != nil && t != nil {
		if t.Type == lexer.Star || t.Type == lexer.Slash {
			t = r.Read()
			child2 := c.primary()
			if child2 != nil {
				n = &parser.ASTNode{Type: parser.Multiplicative, Text: t.Text}
				n.AddChild(child1)
				n.AddChild(child2)
			} else {
				log.Fatalln("invalid multiply expression, expecting the right part.")
			}
		}
	}

	return n
}

//语法解析：基础表达式
func (c *Calculator) primary() *parser.ASTNode {
	r := c.Lexer.Reader
	var n *parser.ASTNode
	t := r.Peek()
	if t != nil {
		if t.Type == lexer.IntLiteral {
			t = r.Read()
			n = &parser.ASTNode{Type: parser.IntLiteral, Text: t.Text}
		} else if t.Type == lexer.Identifier {
			t = r.Read()
			n = &parser.ASTNode{Type: parser.Identifier, Text: t.Text}
		} else if t.Type == lexer.LeftParen {
			r.Read()
			n = c.add()
			if n != nil {
				t = r.Peek()
				if t != nil && t.Type == lexer.RightParen {
					r.Read()
				} else {
					log.Fatalln("expecting right parenthesis.")
				}
			} else {
				log.Fatalln("expecting an additive expression inside parenthesis")
			}
		}
	}
	return n //简化AST：不构造primary节点，直接返回子节点，只有一个子节点。
}

//打印AST的树结构
// indent 缩进字符，由tab组成，每一级多一个tab
func (c *Calculator) DumpAST(node parser.ASTNode, indent string) {
	fmt.Println(indent + parser.ASTNodeTypeNameMap[node.Type] + " " + node.Text)
	for _, value := range node.Children {
		c.DumpAST(*value, "\t")
	}
}
