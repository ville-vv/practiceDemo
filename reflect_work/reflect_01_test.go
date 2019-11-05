package reflect_work

import (
	"testing"
)

func TestPrintStructInfo(t *testing.T) {
	stud := &Student{Age: 20, Name: "you are bc", Lsn: &Lesson{lTp: &LessonType{}}}
	PrintObjectInfo(stud)
}

func TestStructMethodCall(t *testing.T) {
	stud := &Student{Age: 20, Name: "you are bc", Lsn: &Lesson{lTp: &LessonType{}}}
	StructMethodCall(stud)
}
