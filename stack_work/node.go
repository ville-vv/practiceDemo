// @File     : node
// @Author   : Ville
// @Time     : 19-10-11 下午4:41
// stack
package stack

import "strconv"

type node struct {
	value    interface{}
	previous *node
	Next     *node
}

func (sel *node) ToString() string {
	return ""
}
func (sel *node) ToInt() int {
	d := 0
	switch sel.value.(type) {
	case string:
		d, _ := strconv.Atoi(sel.value.(string))
		return d
	case int64:
		d = int(sel.value.(int64))
	case int32:
		d = int(sel.value.(int32))
	case int:
		d = sel.value.(int)
	}
	return d
}
func (sel *node) ToInt64() int64 {
	d := int64(0)
	switch sel.value.(type) {
	case string:
		d, _ := strconv.ParseInt(sel.value.(string), 10, 64)
		return d
	case int64:
		d = sel.value.(int64)
	case int32:
		d = int64(sel.value.(int32))
	case int:
		d = int64(sel.value.(int))
	}
	return d
}
func (sel *node) ToInt32() int {
	return 0
}
func (sel *node) Value() interface{} {
	return sel.value
}
