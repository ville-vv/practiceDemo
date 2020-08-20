package svc

import (
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io"
	"strings"
)

type GenPdf struct {
	width        float64
	height       float64
	leftMargin   float64
	rightMargin  float64
	topMargin    float64
	bottomMargin float64
	opt          *PdfOpt
	fpdf         *gofpdf.Fpdf
}

func NewGenPdf(opt *PdfOpt) *GenPdf {
	gpdf := &GenPdf{
		opt:          opt,
		leftMargin:   10,
		rightMargin:  10,
		topMargin:    10,
		bottomMargin: 10,
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 11)
	width, height := pdf.GetPageSize()

	gpdf.fpdf = pdf
	gpdf.height = height
	gpdf.width = width
	return gpdf
}

func (g *GenPdf) AddPage() {
	g.fpdf.AddPage()
}

func (g *GenPdf) Do() {
	pdf := g.fpdf
	g.PageFoot()
	pdf.AddPage()

	var kvOpt []*KvOpt
	kvOpt = append(kvOpt, &KvOpt{
		K: NewCellOpt("Our Application/Loan No.").SetBold(true).SetW(0),
		V: NewCellOpt("- {{.OrderNo}}").SetBold(false),
	})

	g.KVShow(&KvShowOpt{list: kvOpt})

	g.TableKv(70, 8, map[string]string{
		"Loan ID":                "Gt20200721003235",
		"Applicant Name":         "Anup Sinha",
		"Email ID":               "aanusinha111@gmail.com",
		"Register Mobile Number": "9891128914",
		"ID Number":              "HILPS5392R",
		"Sanction Date":          "6-Aug-20",
	})

	err := g.fpdf.OutputFileAndClose("textxxx.pdf")
	if err != nil {
		fmt.Println("出库哦呜", err)
		return
	}
}

func (g *GenPdf) Title() {
	g.fpdf.SetTitle("sdfsdfsdfsdf", false)
}

func (g *GenPdf) PageHead() {

}

func (g *GenPdf) PageFoot() {
	pdf := g.fpdf
	oldSize, _ := pdf.GetFontSize()
	pdf.SetFooterFunc(func() {
		pdf.SetFontSize(11)
		pdf.SetY(-15)
		pdf.SetX(-15)
		pdf.CellFormat(0, 10, fmt.Sprintf("%d", pdf.PageNo()), "", 0, "R", false, 0, "")
	})
	pdf.SetFontSize(oldSize)
}

// 文字下划线
func (g *GenPdf) UnderLine() {
}

func (g *GenPdf) H1(x, y float64, isCenter bool, str string) {
	newSize := float64(24)
	g.hx(str, x, y, isCenter, newSize, nil)
}
func (g *GenPdf) H2(x, y float64, isCenter bool, str string) {
	newSize := float64(22)
	g.hx(str, x, y, isCenter, newSize, nil)

}
func (g *GenPdf) H3(x, y float64, isCenter bool, str string) {
	newSize := float64(18)
	g.hx(str, x, y, isCenter, newSize, nil)
}
func (g *GenPdf) H4(x, y float64, isCenter bool, str string) {
	newSize := float64(14)
	g.hx(str, x, y, isCenter, newSize, nil)
}
func (g *GenPdf) H5(x, y float64, isCenter bool, str string) {
	newSize := float64(12)
	g.hx(str, x, y, isCenter, newSize, nil)
}
func (g *GenPdf) H6(x, y float64, isCenter bool, str string) {
	newSize := float64(10)
	g.hx(str, x, y, isCenter, newSize, nil)
}
func (g *GenPdf) H7(x, y float64, isCenter bool, str string) {
	newSize := float64(8)
	g.hx(str, x, y, isCenter, newSize, nil)
}

func (g *GenPdf) hx(str string, x, y float64, isCenter bool, fontSize float64, fc func()) {
	if x == 0 {
		x = g.leftMargin
	}
	pdf := g.fpdf
	align := ""
	strW := pdf.GetStringWidth(str)
	if isCenter {
		x = (g.width - strW) / 2
		fmt.Println(x, g.width, strW)
	}

	newSize := fontSize
	oldSize, _ := pdf.GetFontSize()
	pdf.SetFontSize(newSize)
	pdf.SetFontStyle("B")
	pdf.SetX(x)
	pdf.SetLineWidth(1)
	if fc != nil {
		fc()
	}

	pdf.CellFormat(strW, newSize, str, "0", 0, align, false, 0, "")
	pdf.SetFontSize(oldSize)
	pdf.SetLineWidth(0)
	fmt.Println(oldSize)
	pdf.SetFontStyle("")
	pdf.Ln(newSize * 0.7)
}

// 段落
func (g *GenPdf) P(txtStr string) {
	pdf := g.fpdf
	pdf.CellFormat(g.width, 8, txtStr, "", 0, "", false, 0, "")
	pdf.MultiCell(0, 8, txtStr, "", "L", false)
}

func (g *GenPdf) PB(txtStr string) {
	pdf := g.fpdf
	pdf.SetFontStyle("B")
	pdf.MultiCell(0, 10, txtStr, "", "L", false)
	pdf.SetFontStyle("")
}

func (g *GenPdf) Table(tb *TableOpt) {
	pdf := g.fpdf
	oldFSize, _ := pdf.GetFontSize()
	pdf.SetFontSize(tb.FontSize)

	cellW := func(cOpt *CellOpt) {
		if cOpt.Bold {
			pdf.SetFontStyle("B")
		} else {
			pdf.SetFontStyle("")
		}
		pdf.CellFormat(cOpt.W, cOpt.H, cOpt.Val, "1", 0, cOpt.Align, false, 0, "")
	}
	x := tb.X
	if x == 0 {
		x = g.leftMargin
	}
	for _, row := range tb.Rows {
		pdf.SetX(x)
		for _, cl := range row.Row() {
			cellW(cl)
		}
		pdf.Ln(-1)
	}

	pdf.SetFontSize(oldFSize)
	pdf.SetFontStyle("")

}

func (g *GenPdf) TableKv(w, h float64, vals map[string]string) {
	pdf := g.fpdf
	_, size01 := pdf.GetFontSize()
	pdf.SetFontSize(11)
	for name, val := range vals {
		pdf.SetX(g.leftMargin)
		pdf.SetFontStyle("B")
		pdf.CellFormat(w, h, name, "1", 0, "", false, 0, "")
		pdf.SetFontStyle("")
		pdf.CellFormat(w, h, val, "1", 0, "R", false, 0, "")
		pdf.Ln(-1)
	}
	pdf.SetFontSize(size01)
}

func (g *GenPdf) KVShow(kvs *KvShowOpt) {
	pdf := g.fpdf
	cellW := func(cOpt *CellOpt) {
		if cOpt.Bold {
			pdf.SetFontStyle("B")
		} else {
			pdf.SetFontStyle("")
		}
		if cOpt.W == 0 {
			cOpt.W = pdf.GetStringWidth(cOpt.Val)
		}

		pdf.CellFormat(cOpt.W, cOpt.H, cOpt.Val, "0", 0, cOpt.Align, false, 0, "")
	}

	x := kvs.X
	if x == 0 {
		x = g.leftMargin
	}

	for _, ele := range kvs.list {
		cellW(ele.K)
		cellW(ele.V)
		pdf.Ln(-1)
	}
}

func (g *GenPdf) Image(name string, x, y, w, h float64) {
	//g.fpdf.Image(name, x, y, w, h, false, "", 0, "")

	err := g.ImageReader(x, y, w, h, name, bytes.NewBuffer(MountSkikharImgData))
	if err != nil {
		return
	}

	g.fpdf.Ln(h)
}

func (g *GenPdf) ImageReader(x, y, w, h float64, name string, reader io.Reader) error {
	opt := gofpdf.ImageOptions{
		ReadDpi: false,
	}
	pos := strings.LastIndex(name, ".")
	if pos < 0 {
		return fmt.Errorf("image file has no extension and no type was specified: %s", name)
	}
	opt.ImageType = name[pos+1:]
	g.fpdf.RegisterImageOptionsReader(name, opt, reader)
	g.fpdf.Image(name, x, y, w, h, false, "", 0, "")
	return nil
}

func (g *GenPdf) Ln(h float64) {
	g.fpdf.Ln(h)
}

func (g *GenPdf) PDF() *gofpdf.Fpdf {
	return g.fpdf
}
