package svc

import (
	"testing"
)

func TestNewGenPdf(t *testing.T) {
	pdf := NewGenPdf(nil)
	pdf.Do()
}
