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

