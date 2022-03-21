package asii85

// It is go implementation of ASCII85/ASCIIHex decoder (Adobe version).
// In ASCII85 encoding, every four bytes are encoded with five ASCII
// letters, using 85 different types of characters (as 256**4 < 85**5).
// When the length of the original bytes is not a multiple of 4, a special
// rule is used for round up.
// The Adobe's ASCII85 implementation is slightly different from
// its original in handling the last characters.
func ascii85_decode(data []byte) {

}

// ASCIIHexDecode filter: PDFReference v1.4 section 3.3.1
// For each pair of ASCII hexadecimal digits (0-9 and A-F or a-f), the
// ASCIIHexDecode filter produces one byte of binary data. All white-space
// characters are ignored. A right angle bracket character (>) indicates
// EOD. Any other characters will cause an error. If the filter encounters
// the EOD marker after reading an odd number of hexadecimal digits, it
// will behave as if a 0 followed the last digit.
func ascii85_encode(data []byte) {

}
