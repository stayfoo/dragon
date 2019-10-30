/*
@Time : 2019-10-28 13:12
@Author : mengyueping
@File : ast_node_type
@Software: GoLand
*/
package parser

//AST节点的类型
type ASTNodeType int

const (
	Program ASTNodeType = iota //程序入口，根节点

	IntDeclaration //整型变量声明
	ExpressionStmt //表达式语句，即表达式后面跟个分号
	AssignmentStmt //赋值语句

	Primary        //一元表达式
	Multiplicative //乘法表达式，除法表达式
	Additive       //加法表达式，减法表达式

	Identifier //标识符
	IntLiteral //整型字面量
)

var ASTNodeTypeNameMap = map[ASTNodeType]string{
	Program: "Program",

	IntDeclaration: "IntDeclaration",
	ExpressionStmt: "ExpressionStmt",
	AssignmentStmt: "AssignmentStmt",

	Primary:        "Primary",
	Multiplicative: "Multiplicative",
	Additive:       "Additive",

	Identifier: "Identifier",
	IntLiteral: "IntLiteral",
}
