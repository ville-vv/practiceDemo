package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

var (
	String string
	Number int
	Input  string
)

func getInput() {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		fmt.Print("请输入用户名-> ")
		Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
		//if len(Input) == 1 {
		//	continue //如果用户输入的是一个空行就让用户继续输入。
		//}
		fmt.Printf("您输入的是:%s\n", Input)
		if String == "stop" {
			break
		}
		fmt.Printf("输入密码-> ")
		pas, _ := GetPass(os.Stdin, os.Stdout)
		fmt.Printf("您输入的密码是:%s \n", string(pas))
	}
}

var getCh = func(r io.Reader) (byte, error) {
	buf := make([]byte, 1)
	if n, err := r.Read(buf); n == 0 || err != nil {
		if err != nil {
			return 0, err
		}
		return 0, io.EOF
	}
	return buf[0], nil
}

// 获取密码
func GetPass(r *os.File, w io.Writer) ([]byte, error) {
	var err error
	var pass, bs []byte
	//主要就是这一行终端文件替换
	if terminal.IsTerminal(int(r.Fd())) {
		oldState, err := terminal.MakeRaw(int(r.Fd()))
		if err != nil {
			return pass, err
		}
		defer func() {
			terminal.Restore(int(r.Fd()), oldState)
			fmt.Println("")
		}()
	}

	var counter int
	for counter = 0; counter <= 200; counter++ {
		if v, e := getCh(r); e != nil {
			err = e
			break
		} else if v == 127 || v == 8 {
			if l := len(pass); l > 0 {
				pass = pass[:l-1]
				fmt.Fprint(w, string(bs))
			}
		} else if v == 13 || v == 10 {
			break
		} else if v == 3 {
			err = errors.New("interrupted")
			break
		} else if v != 0 {
			pass = append(pass, v)
			fmt.Fprint(w, "*")
		}
	}
	return pass, err

}

func main() {
	getInput()
}
