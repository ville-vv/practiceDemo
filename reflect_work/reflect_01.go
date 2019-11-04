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

func (s *Student) Playing() {
	fmt.Println("I like to playing basketball")
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
