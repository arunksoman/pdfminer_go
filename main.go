package main

import (
	"encoding/hex"
	"fmt"

	"github.com/arunksoman/pdfminer_go/arcfour"
)

func main() {
	a := arcfour.ArcFourEncrypt("Wiki", "pedia")
	fmt.Println(hex.EncodeToString(a))
}
