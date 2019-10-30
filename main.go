/*
@Time : 2019-08-25 12:15
@Author : mengyueping
@File : main
@Software: GoLand
*/
package main

import (
	"bufio"
	"dragon/calculator"
	"dragon/script"
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

var code string
var verbose string
func init() {
	flag.StringVar(&code, "code", "", "Run your script.")
	flag.StringVar(&verbose, "verbose", "", "Verbose the process of script execution. It will show AST and more info.")
}

func main() {
	flag.Parse()
	//fmt.Println("code: ", code)

	s := script.GoScript{}
	args := []string{}
	s.REPL(args, false)

	//Example()
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
	//code13 := "2+3+;"
	//calculatorTest(code13)
	//code14 := "2+3*;"
	//calculatorTest(code14)

}

func calculatorTest(code string) {
	cal := calculator.Calculator{}
	cal.Code = code
	cal.Evaluate()
}

//获取标准输入
func scanInput()  {
	fmt.Println("----方法1,输入：---")
	//方法1：
	rd := bufio.NewReader(os.Stdin)
	input, err := rd.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)


	//方法2：
	fmt.Println("----方法2,输入：---")
	rd2 := bufio.NewReader(os.Stdin)
	line, _, err := rd2.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(line))


	//方法3：
	fmt.Println("----方法3,输入：---")
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		text := scan.Text()
		fmt.Println(text)
	}
}

func demo() {
	str := "This is golang."
	runeStr := []rune(str)
	fmt.Println("runeStr: ", runeStr)

	fmt.Println("str len: ", len(str))                            //15
	fmt.Println("rune len: ", len(runeStr))                       //15
	fmt.Println("rune string len: ", utf8.RuneCountInString(str)) //15

	for _, value := range runeStr {
		fmt.Println("value: ", value)
		fmt.Println("value str: ", string(value))
	}

	//会更改原来切片的数据
	copy(runeStr[0:3], []rune("it"))

	fmt.Println("str: ", str)
	fmt.Println("runeStr: ", runeStr)

	cStr := string(runeStr)
	fmt.Println("cStr: ", cStr)

	//lib.Test()
	//ReverseString(str)
	Tooargs(1, 2, 3, 4, 5, 6, 7, 8)
}

//变参数
func Tooargs(args ...int) {
	for key, value := range args {
		fmt.Println("key: ", key, "value: ", value)
	}
}

func ReverseString(str string) {
	strRune := []rune(str)
	result := ""
	len := len(strRune)
	for j := len - 1; j >= 0; j-- {
		item := strRune[j]
		result = result + string(item)
	}
	fmt.Println(result)
}

func DemoGoto() {
	i := 0
	fmt.Println("----- i: ", i)
MARK:
	for i < 10 {
		i++
		fmt.Println("i: ", i)
		goto MARK
	}
}
