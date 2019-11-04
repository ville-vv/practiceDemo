package reflect_work

import (
	"encoding/json"
	"testing"
)

func TestPrintStructInfo(t *testing.T) {
	stud := &Student{Age: 20, Name: "you are bc", Lsn: &Lesson{lTp: &LessonType{}}}
	PrintObjectInfo(stud)
	json.Unmarshal([]byte("{}"), stud)
}
