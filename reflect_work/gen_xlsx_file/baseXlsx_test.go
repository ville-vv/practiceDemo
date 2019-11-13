package xlsxr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

type XlslRowStruct struct {
	Amount     float64    `title:"口袋的钱"`
	FirstName  string     `title:"姓名"`
	Age        uint32     `title:"年龄"`
	Phone      string     `title:"手机号"`
	Card       string     `title:"卡号"`
	Occupation string     `title:"职业"`
	Address    string     `title:"地址"`
	Birthday   *time.Time `title:"开始时间" format:"2006/01/02"`
	Remark     string     `title:"备注"`
	BKD        string     `title:"进欧冠呢"`
}

func TestDataToXlsx_ToXlsx(t *testing.T) {
	x := NewMakeXlsx()
	tm := time.Now()
	table1 := &XlsxTable{
		SheetName: "第2个标线",
		Bodys: []*XlslRowStruct{
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 24, Phone: "1234567890", Card: "PIFG0988X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 14, Phone: "1234577890", Card: "PIFG0989X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 35, Phone: "1234587890", Card: "PIFG0981X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 36, Phone: "1234597890", Card: "PIFG0982X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 44, Phone: "1234507890", Card: "PIFG0983X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 38, Phone: "1234517890", Card: "PIFG0984X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 31, Phone: "1234527890", Card: "PIFG0985X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 30, Phone: "1234537890", Card: "PIFG0986X", Remark: "圣诞节佛鞥欧迪芬金额就感觉", BKD: "大家佛开始的感觉欧克"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Json", Age: 32, Phone: "1234547890", Card: "PIFG0987X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
		},
	}

	table2 := &XlsxTable{
		SheetName: "第一个标线",
		Bodys: []*XlslRowStruct{
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 24, Phone: "1234567890", Card: "PIFG0988X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 14, Phone: "1234577890", Card: "PIFG0989X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 35, Phone: "1234587890", Card: "PIFG0981X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 36, Phone: "1234597890", Card: "PIFG0982X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 44, Phone: "1234507890", Card: "PIFG0983X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 38, Phone: "1234517890", Card: "PIFG0984X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 31, Phone: "1234527890", Card: "PIFG0985X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 30, Phone: "1234537890", Card: "PIFG0986X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
			{Birthday: &tm, Amount: 2349.01, FirstName: "Kson", Age: 32, Phone: "1234547890", Card: "PIFG0987X", Remark: "圣诞节佛鞥欧迪芬金额就感觉"},
		},
	}
	table3 := &XlsxTable{
		SheetName: "第三个标线",
		Bodys:     []*XlslRowStruct{},
	}

	//buf := bytes.NewBufferString("")
	file, err := os.Create("test.xlsx")
	if err != nil {
		return
	}
	x.ToXlsx([]*XlsxTable{table1, table2, table3}, file)
}

func TestTitle(t *testing.T) {
	titles, err := Title(&XlslRowStruct{})
	assert.NoError(t, err)
	fmt.Println(titles)

	titles, err = Title([]*XlslRowStruct{})
	assert.NoError(t, err)
	fmt.Println(titles)

	titles, err = Title(map[string]*XlslRowStruct{})
	assert.NoError(t, err)
	fmt.Println(titles)
}
