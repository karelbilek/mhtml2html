package mhtml2html

import "testing"

func TestConvertMht2HTML(t *testing.T) {
	resp, err := ConvertMht2HTML("13.mht")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
