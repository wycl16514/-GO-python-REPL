package repl 

import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) //从控制台获取输入

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan() //点击回车后返回输入内容
		if !scanned {
			return //没有输入内容
		}

		line := scanner.Text() //当前输入的内容
		l := lexer.New(line) 

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok) //输出解析的结果
		}
	}
}