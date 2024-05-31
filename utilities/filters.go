package utilities

import (
	"bytes"
	"log"
	"math"
)

// section 6.6 http://www.libpng.org/pub/png/spec/1.2/PNG-Filters.html

func PaethPredictor(left float64, above float64, upper_left float64) float64 {
	p := left + above - upper_left
	// distance to a, b, c
	pa := math.Abs(p - left)
	pb := math.Abs(p - above)
	pc := math.Abs(p - upper_left)
	// return nearest of a, b, c breaking ties in order a, b, c
	if (pa <= pb) && (pa <= pc) {
		return left
	} else if pb <= pc {
		return above
	} else {
		return upper_left
	}
}

func ApplyPngPredictor(pred int, colors int, columns int, bitspercomponet int, data []byte) []byte {
	/*
		Reverse The Effect of Filter

		Documentation: http://www.libpng.org/pub/png/spec/1.2/PNG-Filters.html
	*/
	if bitspercomponet != 8 && bitspercomponet != 1 {
		log.Fatalf("Unsupported `bitspervcomponent`: %d", bitspercomponet)
	}
	nbytes := colors * columns * int(math.Floor(float64(bitspercomponet)/8))
	bpp := colors * int(math.Floor(float64(bitspercomponet)/8))
	buf := []byte{}
	line_above := bytes.Repeat([]byte{0x00}, columns)
	for scanline_i := 0; scanline_i < len(data); scanline_i += nbytes + 1 {
		filter_type := data[scanline_i]
		line_encoded := data[scanline_i+1 : scanline_i+1+nbytes]
		raw := []byte{}
		if filter_type == 0 {
			// Filter type 0: None
			raw = append(raw, line_encoded...)
		} else if filter_type == 1 {
			// Filter type 1: Sub
			// To reverse the effect of the Sub() filter after decompression,
			// output the following value:
			//   Raw(x) = Sub(x) + Raw(x - bpp)
			// (computed mod 256), where Raw() refers to the bytes already
			//  decoded.
			for j, sub_x := range line_encoded {
				var raw_x_bpp byte
				if j-bpp < 0 {
					raw_x_bpp = 0
				} else {
					raw_x_bpp = raw[j-bpp]
				}
				raw_x := (sub_x + raw_x_bpp) & 255
				raw = append(raw, raw_x)
			}
		} else if filter_type == 2 {
			// Filter type 2: Up
			// To reverse the effect of the Up() filter after decompression,
			// output the following value:
			//   Raw(x) = Up(x) + Prior(x)
			// (computed mod 256), where Prior() refers to the decoded bytes of
			// the prior scanline.
			var l_length int
			if len(line_encoded) > len(line_above) {
				l_length = len(line_above)
			} else if len(line_above) > len(line_encoded) {
				l_length = len(line_encoded)
			} else {
				l_length = len(line_encoded)
			}
			for i := 0; i < l_length; i++ {
				up_x := line_encoded[i]
				prior_x := line_above[i]
				raw_x := (up_x + prior_x) & 255
				raw = append(raw, raw_x)
			}
		} else if filter_type == 3 {
			// Filter type 3: Average
			// To reverse the effect of the Average() filter after
			// decompression, output the following value:
			//    Raw(x) = Average(x) + floor((Raw(x-bpp)+Prior(x))/2)
			// where the result is computed mod 256, but the prediction is
			// calculated in the same way as for encoding. Raw() refers to the
			// bytes already decoded, and Prior() refers to the decoded bytes of
			// the prior scanline.
			for j, average_x := range line_encoded {
				var raw_x_bpp byte
				if j-bpp < 0 {
					raw_x_bpp = 0
				} else {
					raw_x_bpp = raw[j-bpp]
				}
				prior_x := raw[j-bpp]
				// raw_x := (average_x + byte(math.Floor(float64((raw_x_bpp+prior_x)/2)))) & 255
				raw_x := (average_x + byte((raw_x_bpp+prior_x)/2)) & 255
				raw = append(raw, raw_x)
			}
		} else if filter_type == 4 {
			// Filter type 4: Paeth
			// To reverse the effect of the Paeth() filter after decompression,
			// output the following value:
			//    Raw(x) = Paeth(x)
			//             + PaethPredictor(Raw(x-bpp), Prior(x), Prior(x-bpp))
			// (computed mod 256), where Raw() and Prior() refer to bytes
			// already decoded. Exactly the same PaethPredictor() function is
			// used by both encoder and decoder.
			for j, paeth_x := range line_encoded {
				var raw_x_bpp byte
				var prior_x_bpp byte
				if j-bpp < 0 {
					raw_x_bpp = 0
					prior_x_bpp = 0
				} else {
					raw_x_bpp = raw[j-bpp]
					prior_x_bpp = line_above[j-bpp]
				}
				prior_x := line_above[j]
				paeth := PaethPredictor(float64(raw_x_bpp), float64(prior_x), float64(prior_x_bpp))
				raw_x := (paeth_x + byte(paeth)) & 255
				raw = append(raw, raw_x)
			}
		} else {
			log.Fatalf("Unsupported predictor value: %d", filter_type)
		}
		buf = append(buf, raw...)
		line_above = raw
	}
	return buf
}
