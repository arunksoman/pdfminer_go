package ascii85

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// It is go implementation of ASCII85/ASCIIHex decoder (Adobe version).
// In ASCII85 encoding, every four bytes are encoded with five ASCII
// letters, using 85 different types of characters (as 256**4 < 85**5).
// When the length of the original bytes is not a multiple of 4, a special
// rule is used for round up.
// The Adobe's ASCII85 implementation is slightly different from
// its original in handling the last characters.
func Ascii85_Decode(data []byte) []byte {
	n := 0
	var b uint32 = 0
	var out []byte
	buf := new(bytes.Buffer)
	for i := 0; i < len(data); i++ { // 33 <= ascii85 <=117
		if '!' <= data[i] && data[i] <= 'u' {
			n += 1
			b = b*85 + (uint32(data[i]) - 33)
			if n == 5 {
				err := binary.Write(buf, binary.BigEndian, b)
				if err != nil {
					fmt.Println("binary.Write failed:", err)
				}
				for n_bytes := 0; n_bytes < len(buf.Bytes()); n_bytes++ {
					out = append(out, buf.Bytes()[n_bytes])
				}

				n = 0
				b = 0
			}
		} else if data[i] == 'z' { // Have to check. As of now, no idea
			zeros := []byte{0x00, 0x00, 0x00, 0x00}
			if n == 0 {
				for j := 0; j < 4; j++ {
					out = append(out, zeros[j])
				}
				n = 0
			}
		} else if data[i] == '~' { // end of data
			if n != 0 {
				for j := 0; j < (5 - n); j++ {
					b = b*85 + 84
				}
				err := binary.Write(buf, binary.BigEndian, b)
				if err != nil {
					fmt.Println("binary.Write failed:", err)
				}
				for n_bytes := 0; n_bytes < len(buf.Bytes())-(n+1); n_bytes++ {
					out = append(out, buf.Bytes()[n_bytes])
				}
			}
			break
		}
		buf.Reset()
	}
	return out
}

// ASCIIHexDecode filter: PDFReference v1.4 section 3.3.1
// For each pair of ASCII hexadecimal digits (0-9 and A-F or a-f), the
// ASCIIHexDecode filter produces one byte of binary data. All white-space
// characters are ignored. A right angle bracket character (>) indicates
// EOD. Any other characters will cause an error. If the filter encounters
// the EOD marker after reading an odd number of hexadecimal digits, it
// will behave as if a 0 followed the last digit.

// func decode(x byte) byte {
// 	i := int16(x)
// 	return byte(i)
// }

// func Ascii85_HexDecode(data []byte) {
// 	hex_re := regexp.MustCompile(`([a-f\d]{2})`)
// 	trail_re := re.MustCompile(`^(?:[a-f\d]{2}|\s)*([a-f\d])[\s>]*$`)
// 	var out []byte
// 	x := hex_re.FindAll(data)
// 	for i := 0; i < len(data); i++ {

// 	}
// }
