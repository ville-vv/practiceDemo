package xlsxr

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

// json 转 bytes 特殊字符适用
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)

	var out bytes.Buffer
	if err = json.Indent(&out, buffer.Bytes(), "", "\t"); err != nil {
		return nil, err
	}
	return out.Bytes(), err
}

// 把 struct 保存为已格式化的json文件
func SaveFileForJson(fileName string, obj interface{}) (err error) {
	fileInfos := strings.Split(fileName, ".")
	if fileInfos[len(fileInfos)-1] != "json" {
		fileName += ".json"
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	buf, err := JSONMarshal(obj)
	if err != nil {
		return nil
	}
	_, err = file.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

type XlsxTable struct {
	SheetName string
	Bodys     interface{}
}

func (sel *XlsxTable) getType(val reflect.Type) reflect.Type {
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Map {
		val = val.Elem()
	}
	switch val.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		val = val.Elem()
	default:
		panic("type is error")
	}
	return val
}

func (sel *XlsxTable) Title() []string {
	refl := sel.getType(reflect.TypeOf(sel.Bodys))
	titles := make([]string, refl.NumField())
	for i := 0; i < refl.NumField(); i++ {
		titles[i] = refl.Field(i).Tag.Get("title")
	}
	return titles
}

func (sel *XlsxTable) makeTitle(sheet *xlsx.Sheet) {
	var row *xlsx.Row
	if len(sheet.Rows) > 0 {
		row = sheet.Rows[0]
	} else {
		row = sheet.AddRow()
	}
	titles := sel.Title()
	for _, v := range titles {
		cell := row.AddCell()
		cell.SetString(v)
	}
}

func (sel *XlsxTable) makeBody(sheet *xlsx.Sheet) {

	datas := reflect.ValueOf(sel.Bodys)

	num := datas.Len()

	for i := 0; i < num; i++ {

		refTr := datas.Index(i)
		if refTr.Kind() == reflect.Ptr {
			refTr = refTr.Elem()
		}
		row := sheet.AddRow()
		for i := 0; i < refTr.NumField(); i++ {
			cell := row.AddCell()
			if strings.Contains(refTr.Field(i).Type().String(), "time.Time") {
				cell.SetString(sel.timeFormat(refTr.Field(i), refTr.Type().Field(i).Tag.Get("format")))
				continue
			}
			cell.SetValue(refTr.Field(i).Interface())
		}
	}
}

func (sel *XlsxTable) timeFormat(valV reflect.Value, format string) string {
	if format != "" {
		itm := valV.Interface()
		if valV.Kind() == reflect.Ptr {
			return itm.(*time.Time).Format(format)
		}
		return itm.(time.Time).Format(format)
	}
	return valV.String()
}

func (sel *XlsxTable) ToSheet() *xlsx.Sheet {
	sheet := &xlsx.Sheet{Name: sel.SheetName}
	sel.makeTitle(sheet)
	sel.makeBody(sheet)
	return sheet
}

func (sel *XlsxTable) ToXlsx(writer io.Writer) {
	xlsxFile := xlsx.NewFile()
	sheet := sel.ToSheet()
	xlsxFile.AppendSheet(*sheet, sel.SheetName)
	xlsxFile.Write(writer)
	return
}

type MakeXlsx struct {
	xlsxFile *xlsx.File
	xsheets  map[string]*xlsx.Sheet
}

func NewMakeXlsx() *MakeXlsx {
	return &MakeXlsx{
		xlsxFile: xlsx.NewFile(),
		xsheets:  make(map[string]*xlsx.Sheet),
	}
}

func (sel *MakeXlsx) addSheet(name string) (*xlsx.Sheet, error) {
	if sheet, ok := sel.xsheets[name]; ok {
		return sheet, nil
	}
	sheet, err := sel.xlsxFile.AddSheet(name)
	if err != nil {
		return nil, err
	}
	sel.xsheets[name] = sheet
	return sheet, nil
}

func (sel *MakeXlsx) makeTitle(sheet *xlsx.Sheet, titles []string) {
	var row *xlsx.Row
	if len(sheet.Rows) > 0 {
		row = sheet.Rows[0]
	} else {
		row = sheet.AddRow()
	}
	for _, v := range titles {
		row.AddCell().SetString(v)
	}
}

func (sel *MakeXlsx) makeBody(sheet *xlsx.Sheet, bodys interface{}) {
	datas := reflect.ValueOf(bodys)
	num := datas.Len()
	for i := 0; i < num; i++ {
		refTr := datas.Index(i)
		if refTr.Kind() == reflect.Ptr {
			refTr = refTr.Elem()
		}
		row := sheet.AddRow()
		for i := 0; i < refTr.NumField(); i++ {
			if strings.Contains(refTr.Field(i).Type().String(), "time.Time") {
				row.AddCell().SetString(sel.timeFormat(refTr.Field(i), refTr.Type().Field(i).Tag.Get("format")))
				continue
			}
			row.AddCell().SetValue(refTr.Field(i).Interface())
		}
	}
}

func (sel *MakeXlsx) timeFormat(valV reflect.Value, format string) string {
	if valV.IsZero() {
		return ""
	}
	if format != "" {
		itm := valV.Interface()
		if valV.Kind() == reflect.Ptr {
			return itm.(*time.Time).Format(format)
		}
		return itm.(time.Time).Format(format)
	}
	return valV.String()
}

func (sel *MakeXlsx) ToXlsx(datas []*XlsxTable, wt io.Writer) error {
	xlsxFile := sel.xlsxFile
	for _, xd := range datas {
		sheet, err := sel.addSheet(xd.SheetName)
		if err != nil {
			return err
		}
		sel.makeTitle(sheet, xd.Title())
		sel.makeBody(sheet, xd.Bodys)
	}
	xlsxFile.Write(wt)
	return nil
}

func Title(as interface{}) ([]string, error) {
	val, err := getType(reflect.TypeOf(as))
	if err != nil {
		return nil, err
	}
	num := val.NumField()
	titles := make([]string, num)
	for i := 0; i < num; i++ {
		titles[i] = val.Field(i).Tag.Get("title")
	}
	return titles, nil
}

func getType(val reflect.Type) (reflect.Type, error) {
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Map {
		val = val.Elem()
	}
	switch val.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		val = val.Elem()
	default:
		return nil, errors.New("type is error")
	}
	return val, nil
}

func typeCheck(val reflect.Type) (reflect.Type, error) {
	switch val.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		val = val.Elem()
	case reflect.Slice:
		return typeCheck(val.Elem())
	default:
		return nil, errors.New("type is error")
	}
	return val, nil
}
