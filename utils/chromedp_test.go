package utils

import (
	"os"
	"testing"
)

func TestPrintToPdf(t *testing.T) {

	cookies := map[string]string{
		"GCESS":    "BAoEAAAAAAsCBAAMAQECBEi0pl4GBN1q2sgFBAAAAAAEBAAvDQAIAQMJAQEBBOjmFAAHBEwyTEsDBEi0pl4-",
		"GCID":     "d2bd332-046e3e7-b5611b2-e95032b",
		"GRID":     "d2bd332-046e3e7-b5611b2-e95032b",
		"SERVERID": "1fa1f330efedec1559b3abbcb6e30f50|1587983432|1587983432",
	}

	filename := "file.pdf"
	err := PrintToPDF(185527, filename, cookies)

	if err != nil {
		t.Fatal("PrintToPDF test is failure", err)
	} else {
		os.Remove(filename)
	}
}
