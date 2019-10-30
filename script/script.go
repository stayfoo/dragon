/*
@Time : 2019-10-29 19:41
@Author : mengyueping
@File : script
@Software: GoLand
*/
package script

import (
	"bufio"
	"dragon/calculator"
	"fmt"
	"io"
	"strings"
	"testing/iotest"

	"os"
	//"strings"
	//"testing/iotest"
)

/**
一个简单的脚本解释器。
所支持的语法，请参见parser.go

运行脚本：
在命令行下，键入：go script
则进入一个REPL界面。你可以依次敲入命令。比如：
 > 2+3;
 > int age = 10;
 > int b;
 > b = 10*2;
 > age = age + b;
 > exit();  //退出REPL界面。

你还可以使用一个参数 -v，让每次执行脚本的时候，都输出AST和整个计算过程。
*/

type GoScript struct {
	Variables map[string]int
}

//实现一个 REPL
func (s *GoScript) REPL(args []string, verbose bool) {
	fmt.Println("It is a script language. It is written by golang!\nInput your code.")
	var scriptText string
	rd := bufio.NewReader(os.Stdin) //读取输入的内容
	println("\n>") //提示符

	for {
		input, err := rd.ReadString('\n') //定义一行输入的内容分隔符
		if err == io.EOF {
			break
		}
		if err != nil && err != iotest.ErrTimeout {
			panic("GetLines: " + err.Error())
		}

		if verbose {
			fmt.Println("you input: ", input)
		}

		lineText := strings.TrimSpace(input)
		if lineText == "exit();" {
			fmt.Println("codding end!")
			break
		}

		scriptText += lineText + "\n"
		if strings.HasSuffix(lineText, ";") {
			s.evaluate(scriptText, verbose)

			fmt.Println("\n>") //提示符
			scriptText = ""
		}
	}
}

//计算值
func (s *GoScript) evaluate(code string, verbose bool) int {

	cal := calculator.Calculator{}
	cal.Verbose = verbose
	cal.Code = code
	if len(s.Variables) > 0 {
		cal.Variables = s.Variables
	}
	cal.Evaluate()

	if len(cal.Variables) > 0 {
		for key, value := range cal.Variables {
			s.Variables[key] = value
		}
	}
	//fmt.Println("s.Variables: ", s.Variables)
	//fmt.Println("cal.Variables: ", cal.Variables)

	fmt.Println("result: ", fmt.Sprintf("%d",cal.Result))

	return cal.Result
}