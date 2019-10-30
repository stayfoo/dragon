/*
@Time : 2019-10-28 13:12
@Author : mengyueping
@File : ast_node
@Software: GoLand
*/
package parser

import "fmt"

//AST 节点
//属性包括：AST的类型、文本值、下级子节点和父节点
type ASTNode struct {
	//父节点
	Parent *ASTNode
	//子节点
	Children []*ASTNode
	//文本值
	Text string
	//AST类型
	Type ASTNodeType
}

//添加子节点
func (n *ASTNode) AddChild(node *ASTNode) {
	n.Children = append(n.Children, node)
	node.Parent = n
}

//打印输出AST的树状结构
//indent 缩进字符，由tab组成，每一级多一个tab
func (n *ASTNode)DumpAST(indent string) {
	fmt.Println(indent + ASTNodeTypeNameMap[n.Type] + " " + n.Text)
	for _, value := range n.Children {
		value.DumpAST(indent+"\t")
	}
}
