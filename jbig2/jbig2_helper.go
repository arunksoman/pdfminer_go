package jbig2

import (
	"encoding/binary"
	"errors"
)

var SEG_STRUCT = []struct {
	format string
	name   string
}{
	{">L", "number"},
	{">B", "flags"},
	{">B", "retention_flags"},
	{">B", "page_assoc"},
	{">L", "data_length"},
}

// segment header literals
const HEADER_FLAG_DEFERRED = 0b10000000
const HEADER_FLAG_PAGE_ASSOC_LONG = 0b01000000

const SEG_TYPE_MASK = 0b00111111

const REF_COUNT_SHORT_MASK = 0b11100000
const REF_COUNT_LONG_MASK = 0x1FFFFFFF
const REF_COUNT_LONG = 7

const DATA_LEN_UNKNOWN = 0xFFFFFFFF

// segment types
const SEG_TYPE_IMMEDIATE_GEN_REGION = 38
const SEG_TYPE_END_OF_PAGE = 49
const SEG_TYPE_END_OF_FILE = 51

// file literals
var FILE_HEADER_ID = []byte{0x97, 0x4A, 0x42, 0x32, 0x0D, 0x0A, 0x1A, 0x0A}

const FILE_HEAD_FLAG_SEQUENTIAL = 0b00000001

func bitSet(bit_pos int, value int) bool {
	return (value>>bit_pos)&1 == 1
}

func checkFlag(flag int, value int) bool {
	return flag&value != 0
}

func maskedValue(mask int, value int) (int, error) {
	for bit_pos := 0; bit_pos < 32; bit_pos++ {
		if bitSet(bit_pos, mask) {
			return (value & mask) >> bit_pos, nil
		}
	}
	return 0, errors.New("INVALID `mask` OR `value`")
}

func maskValue(mask int, value int) (int, error) {
	for bit_pos := 0; bit_pos < 32; bit_pos++ {
		if bitSet(bit_pos, mask) {
			return (value & (mask >> bit_pos)) << bit_pos, nil
		}
	}
	return 0, errors.New("INVALID `mask` OR `value`")
}

func unpackInt(format string, buffer []byte) (int, error) {
	switch format {
	case ">B":
		if len(buffer) < 1 {
			return 0, errors.New("buffer too short for uint8")
		}
		return int(buffer[0]), nil
	case ">I", ">L":
		if len(buffer) < 4 {
			return 0, errors.New("buffer too short for uint32")
		}
		result := binary.BigEndian.Uint32(buffer)
		return int(result), nil
	default:
		return 0, errors.New("unsupported format: " + format)
	}
}
