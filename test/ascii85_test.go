package test

import (
	"testing"

	"github.com/arunksoman/pdfminer_go/ascii85"
)

func TestAscii85Decode(t *testing.T) {
	got1 := ascii85.Ascii85_Decode([]byte("9jqo^BlbD-BleB1DJ+*+F(f,q"))
	expected1 := "Man is distinguished"
	if string(got1) != expected1 {
		t.Errorf("got %q, wanted %q", got1, expected1)
	}
	got2 := ascii85.Ascii85_Decode([]byte("E,9)oF*2M7/c~>"))
	expected2 := "pleasure."
	if string(got2) != expected2 {
		t.Errorf("got %q, wanted %q", got2, expected2)
	}
	got3 := ascii85.Ascii85_Decode([]byte(`9jqo^BlbD-BleB1DJ+*+F(f,q/0JhKF<GL>Cj@.4Gp$d7F!,L7@<6@)/0JDEF<G%<+EV:2F!,O<DJ+*.@<*K0@<6L(Df-\0Ec5e;DffZ(EZee.Bl.9pF"AGXBPCsi+DGm>@3BB/F*&OCAfu2/AKYi(DIb:@FD,*)+C]U=@3BN#EcYf8ATD3s@q?d$AftVqCh[NqF<G:8+EV:.+Cf>-FD5W8ARlolDIal(DId<j@<?3r@:F%a+D58'ATD4$Bl@l3De:,-DJs` + "`8ARoFb/0JMK@qB4^F!,R<AKZ&-DfTqBG%G>uD.RTpAKYo'+CT/5+Cei#DII?(E,9)oF*2M7/c~>"))
	expected3 := "Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure."
	if string(got3) != expected3 {
		t.Errorf("got %q, wanted %q", got3, expected3)
	}
}
