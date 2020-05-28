package ast_add_example

import (
	"context"
	"fmt"
)

type Templ struct {
	ch string
}

type DoAction interface {
	Sum(a int, b int) int
	Sub(a int, b int) int
}

func (*Templ) Sum(a int, b int) int {
	fmt.Println("")
	return a + b
}

func (*Templ) Sub(a int, b int) int {
	return a - b
}

func Close() {
}

// Foo 结构体
type Foo struct {
	i int
}

func ADFB() {
}

// Bar 接口
type Bar interface {
	Do(ctx context.Context) error
}

func (*Foo) Do(ctx context.Context) error {
	sub := func(a int, b int) int {
		return a - b
	}
	c := sub(4, 3)
	fmt.Println("结果：", c)
	return nil
}
