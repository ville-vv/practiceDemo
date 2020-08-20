package svc

type CellOpt struct {
	W     float64
	H     float64
	Val   string
	Bold  bool
	Align string
}

func NewCellOpt(val string) *CellOpt {
	return &CellOpt{W: 50, H: 8, Val: val}
}
func (c *CellOpt) Copy(val string) *CellOpt {
	if val == "" {
		val = c.Val
	}
	return &CellOpt{
		W:     c.W,
		H:     c.H,
		Val:   val,
		Bold:  c.Bold,
		Align: c.Align,
	}
}
func (c *CellOpt) SetAlign(align string) *CellOpt {
	c.Align = align
	return c
}
func (c *CellOpt) SetBold(bold bool) *CellOpt {
	c.Bold = bold
	return c
}
func (c *CellOpt) SetW(w float64) *CellOpt {
	c.W = w
	return c
}
func (c *CellOpt) SetH(h float64) *CellOpt {
	c.H = h
	return c
}

type RowOpt struct {
	row []*CellOpt
}

func NewRowOpt(cap int) *RowOpt {
	row := make([]*CellOpt, 0, cap)
	return &RowOpt{row: row}
}
func (r *RowOpt) Row() []*CellOpt {
	return r.row
}

func (r *RowOpt) AddCell(cl *CellOpt) {
	r.row = append(r.row, cl)
}

type TableOpt struct {
	FontSize float64
	X        float64
	Y        float64
	Rows     []*RowOpt
}

type KvOpt struct {
	K *CellOpt
	V *CellOpt
}

type KvShowOpt struct {
	X    float64
	Y    float64
	list []*KvOpt
}

type ImageOpt struct {
}

type PdfOpt struct {
	PageSize string //A4
}
