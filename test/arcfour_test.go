package test

import (
	"encoding/hex"
	"testing"

	"github.com/arunksoman/pdfminer_go/arcfour"
)

func TestArcFour(t *testing.T) {
	got := arcfour.ArcFourEncrypt("Wiki", "pedia")
	go_get := hex.EncodeToString(got)
	expected := "1021bf0420"
	if go_get != expected {
		t.Errorf("got %q, wanted %q", got, expected)
	}
}
