/*
@Time : 2019-10-30 12:34
@Author : mengyueping
@File : example_test
@Software: GoLand
*/
package lexer

func init() {

}

func Example() {
	//语法分析，打印AST树状图
	code1 := "int age = 45;"
	tokenReaderTest(code1)

	code2 := "age == 45;"
	tokenReaderTest(code2)

	code3 := "inta age = 45;"
	tokenReaderTest(code3)

	code4 := "in age = 45;"
	tokenReaderTest(code4)

	code5 := "age >= 45;"
	tokenReaderTest(code5)

	code6 := "age > 45;"
	tokenReaderTest(code6)

	code7 := "45 * 2;"
	tokenReaderTest(code7)

	code8 := "int age = 45+2; age= 20; age+10*2;"
	tokenReaderTest(code8)


	code9 := "2*3+4"
	tokenReaderTest(code9)

	code10 := "1+2*3+4"
	tokenReaderTest(code10)

	code11 := "1 + 2 * 3 + 4"
	tokenReaderTest(code11)

	code12 := "2+3+4"
	tokenReaderTest(code12)

	//计算器：解析异常
	//code13 := "2+3+;"
	//tokenReaderTest(code13)
	//code14 := "2+3*;"
	//tokenReaderTest(code14)

}

func tokenReaderTest(code string) {
	l := Lexer{}
	l.Tokenize(code)
	l.Reader.Dump()
}
