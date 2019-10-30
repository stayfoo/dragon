<p align="center">
  <a href="#">
    <img height="50" src="https://simpleicons.org/icons/go.svg?sanitize=true">
  </a>
</p>

- dragon 使用 golang 编写的一个编译器前端。包含：词法分析器、语法分析器、计算器、脚本语言等。

```
.
├── LICENSE
├── README.md
├── calculator             #计算器
│   ├── calculator.go
│   └── example_test.go
├── example_test.go
├── go.mod
├── lexer                  #词法分析器
│   ├── example_test.go
│   ├── graphic.go
│   ├── lexer.go
│   ├── token.go
│   ├── token_reader.go
│   └── token_type.go
├── main.go  
├── parser                 #语法分析器
│   ├── ast_node.go
│   ├── ast_node_type.go
│   ├── example_test.go
│   └── parser.go
└── script                 #脚本语言
    └── script.go
```

- 词法分析器

lexer 能够为语法分析器、计算器、简单脚本语言生产 Token。

- 语法分析器

parser 能够解析简单的表达式、变量声明和初始化语句、赋值语句。

它支持的语法规则：

```
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
```

```
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
```

- 计算器

语法规则：

```
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
```

存在问题：计算的结合性是有问题的。递归项在右边，会自然的对应右结合。真正需要的是左结合。


- 脚本语言解释器

script 一个脚本解释器。可以进入一个REPL界面。

运行脚本：
在命令行执行命令：

```bash
go run ../dragon
```

语法规则：

```
你可以依次敲入命令。比如：
 > 2+3;
 > exit();  //退出REPL界面。
```

