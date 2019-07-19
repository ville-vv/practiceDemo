// Package other
package other

import "strconv"

type InterfaceA interface {
	AA()
}

type InterfaceB interface {
	BB()
}

type A struct {
	v int
}

func (a *A) AA() {
	a.v += 1
}

type B struct {
	v int
}

func (b *B) BB() {
	b.v += 1
}

func TypeSwitch(v interface{}) {
	switch v.(type) {
	case InterfaceA:
		v.(InterfaceA).AA()
	case InterfaceB:
		v.(InterfaceB).BB()
	}
}

func TypeString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case []byte:
		return string(v.([]byte))
	case int:
		return strconv.Itoa(v.(int))
	}
	return ""
}

func NormalSwitch(a *A) {
	a.AA()
}

func InterfaceSwitch(v interface{}) {
	v.(InterfaceA).AA()
}
