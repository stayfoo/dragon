/*
@Time : 2019-10-30 11:55 
@Author : mengyueping
@File : example_test
@Software: GoLand
*/
package calculator

func init() {

}

func Example() {
	//语法分析，打印AST树状图
	//计算器：+ - * / 运算
	code9 := "2*3+4"
	calculatorTest(code9)

	code10 := "1+2*3+4"
	calculatorTest(code10)

	code11 := "1 + 2 * 3 + 4"
	calculatorTest(code11)

	code12 := "2+3+4"
	calculatorTest(code12)

	//计算器：解析异常
	code13 := "2+3+;"
	calculatorTest(code13)
	code14 := "2+3*;"
	calculatorTest(code14)

}

func calculatorTest(code string) {
	cal := Calculator{}
	cal.Code = code
	cal.Evaluate()
}

