package svc

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewGenPdf(t *testing.T) {
	pdf := NewGenPdf(nil)
	pdf.Do()

	f, err := os.Open("../Mount Shikhar.Png")
	if err != nil {
		return
	}
	defer f.Close()

	f2, err := os.Create("../Mount_Shikhar.txt")
	if err != nil {
		return
	}
	defer f2.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	f2.WriteString("var data = []byte{")

	fmt.Print(len(body), "\n{")
	//for i, v := range body {
	//	fmt.Print(v, ",")
	//	f2.WriteString(fmt.Sprintf("%d,", v))
	//
	//	if i%100 == 0 {
	//		fmt.Println("")
	//		f2.WriteString("\n")
	//	}
	//}
	fmt.Print("}\n")
	f2.WriteString("}")

}
