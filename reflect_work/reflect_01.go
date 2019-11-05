package reflect_work

import (
	"fmt"
	"reflect"
)

type LessonType struct {
	Title string
	Class int
}

type Lesson struct {
	Name  string
	Score int
	lTp   *LessonType
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
	Class string `json:"class"`
	Lsn   *Lesson
}

func (s *Student) Playing(n int, str string, ls *LessonType) {
	fmt.Println("I like to playing basketball")
	fmt.Println("n is :", n)
	fmt.Println("str is :", str)
	fmt.Println("ls.Class is :", ls.Class)
	fmt.Println("ls.Title is :", ls.Title)
}
func (s Student) Playing2(n int, str string, ls *LessonType) {
	fmt.Println("I like to playing basketball")
	fmt.Println("n is :", n)
	fmt.Println("str is :", str)
	fmt.Println("ls.Class is :", ls.Class)
	fmt.Println("ls.Title is :", ls.Title)
}

func PrintObjectInfo(obj interface{}) {
	if obj == nil {
		return
	}
	val := reflect.ValueOf(obj)
	typeD := reflect.TypeOf(obj)
	switch typeD.Kind().String() {
	case "ptr":
		printStructInfo(typeD.Elem(), val.Elem())
	case "struct":
		printStructInfo(typeD, val)
	}
}

func printStructInfo(p reflect.Type, v reflect.Value) {
	if !v.CanAddr() {
		return
	}
	for i := 0; i < p.NumField(); i++ {
		field := p.Field(i)
		switch field.Type.Kind() {
		case reflect.Ptr:
			fmt.Printf("FieldName:%s,\t\t FieldType:%s,\t\t FieldTag:%v \n", field.Name, field.Type, field.Tag)
			// field.PkgPath 为空值是可导出的Field
			if field.PkgPath == "" {
				PrintObjectInfo(v.Field(i).Interface())
			}
		case reflect.Struct:
			PrintObjectInfo(v.Field(i))
		default:
			fmt.Printf("FieldName:%s,\t\t FieldType:%s,\t\t FieldTag:%v,\t\t FieldValue: %v\n", field.Name, field.Type, field.Tag, v.Field(i))
		}
	}
}

func StructMethodCall(v interface{}) {
	val := reflect.ValueOf(v)
	switch val.Kind().String() {
	case "ptr":
		structMethodCall(val)
	case "struct":
		structMethodCall(val)
	}

}

func structMethodCall(v reflect.Value) {
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		fmt.Printf("方法名称：%s \n", v.Type().Method(i).Name)
		args := make([]reflect.Value, 0, 10)
		for j := 0; j < method.Type().NumIn(); j++ {
			in := method.Type().In(j)
			switch in.Kind() {
			case reflect.Ptr:
				fmt.Printf("参数%d：%v, \n", j, in.Elem().Name())
				args = append(args, reflect.ValueOf(&LessonType{Title: "感同身受", Class: 9}))
			case reflect.Int:
				fmt.Printf("参数%d：%v \n", j, in.Name())
				args = append(args, reflect.ValueOf(90))
			case reflect.String:
				fmt.Printf("参数%d：%v \n", j, in.Name())
				args = append(args, reflect.ValueOf("叫几个吗"))
			}
		}
		v.Method(i).Call(args)
	}
}
