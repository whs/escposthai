// escposthai contains Thai print support for ESCPOS printers
// Printer must support codepage no. 20 (Thai Code 42)
// And must be using that mode already. To switch mode:
//
//     p.WriteRaw([]byte{27, 116, 20, 13})

package escposthai

import (
	"github.com/kenshaw/escpos"
)

const space = 32

var cp20 = map[rune]byte{
	'๐': 144,
	'๑': 145,
	'๒': 146,
	'๓': 147,
	'๔': 148,
	'๕': 149,
	'๖': 150,
	'๗': 151,
	'๘': 152,
	'๙': 153,
	'ฃ': 154,
	'ฅ': 155,
	'ก': 161,
	'ข': 162,
	'ค': 163,
	'ฆ': 164,
	'ง': 165,
	'จ': 166,
	'ฉ': 167,
	'ช': 168,
	'ซ': 169,
	'ฌ': 170,
	'ญ': 171,
	'ฎ': 172,
	'ฏ': 173,
	'ฐ': 174,
	'ฑ': 175,
	'ฒ': 176,
	'ณ': 177,
	'ด': 178,
	'ต': 179,
	'ถ': 180,
	'ท': 181,
	'ธ': 182,
	'น': 183,
	'บ': 184,
	'ป': 185,
	'ผ': 186,
	'ฝ': 187,
	'พ': 188,
	'ฟ': 189,
	'ภ': 190,
	'ม': 191,
	'ย': 192,
	'ร': 193,
	'ฤ': 194,
	'ล': 195,
	'ว': 196,
	'ศ': 197,
	'ษ': 198,
	'ส': 199,
	'ห': 200,
	'ฬ': 201,
	'อ': 202,
	'ฮ': 203,
	'ะ': 204,
	'ฦ': 205,
	'า': 206,
	'ำ': 207,
	'เ': 208,
	'แ': 209,
	'โ': 210,
	'ใ': 211,
	'ไ': 212,
	'ๆ': 213,
	'ฯ': 214,
	'ุ': 215,
	'ู': 216,
	'ิ': 217,
	'ี': 218,
	'ึ': 219,
	'ื': 220,
	'ั': 221,
	'ํ': 222,
	'็': 223,
	'่': 224,
	'้': 225,
	'๊': 226,
	'๋': 227,
	'์': 228,
	'ฺ': 229,
}

var mergeMap = map[byte]map[byte]byte{
	222: {
		224: 230, // อ็่
		225: 231, // อ็้
		226: 232, // อ็๊
		227: 233, // อ็๋
	},
	221: {
		224: 234, // อั่
		225: 235, // อั้
		226: 236, // อั๊
		227: 237, // อั๋
	},
	217: {
		224: 238, // อิ่
		225: 239, // อิ้
		226: 240, // อิ๊
		227: 241, // อิ๋
		228: 242, // อิ์
	},
	218: {
		224: 243, // อี่
		225: 244, // อี้
		226: 245, // อี๊
		227: 246, // อี๋
	},
	219: {
		224: 247, // อึ่
		225: 248, // อึ้
		226: 249, // อึ๊
		227: 250, // อึ๋
	},
	220: {
		224: 251, // อื่
		225: 252, // อื้
		226: 253, // อื๊
		227: 254, // อื๋
	},
}

var underChars = []byte{215, 216, 229}
var overChars = []byte{217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228}

func isUnder(r byte) bool {
	for _, b := range underChars {
		if b == r {
			return true
		}
	}
	return false
}

func isOver(r byte) bool {
	for _, b := range overChars {
		if b == r {
			return true
		}
	}
	return false
}

func mergeUpper(existing, with byte) byte {
	emm, ok := mergeMap[existing]
	if !ok {
		return with
	}

	out, ok := emm[with]
	if !ok {
		return with
	}
	return out
}

func scanUpper(text string) (out []byte) {
	for _, ch := range text {
		mapped, ok := cp20[ch]
		if !ok {
			out = append(out, space)
		} else if isOver(mapped) {
			out[len(out)-1] = mergeUpper(out[len(out)-1], mapped)
		} else if isUnder(mapped) {
		} else {
			out = append(out, space)
		}
	}
	out = append(out, byte('\n'))
	return
}

func scanMiddle(text string) (out []byte) {
	for _, ch := range text {
		mapped, ok := cp20[ch]
		if !ok {
			out = append(out, []byte(string(ch))...)
		} else if isOver(mapped) || isUnder(mapped) {
		} else {
			out = append(out, mapped)
		}
	}
	out = append(out, byte('\n'))
	return
}

func scanLower(text string) (out []byte) {
	for _, ch := range text {
		mapped, ok := cp20[ch]
		if !ok {
			out = append(out, space)
		} else if isUnder(mapped) {
			out[len(out)-1] = mapped
		} else if isOver(mapped) {
		} else {
			out = append(out, space)
		}
	}
	out = append(out, byte('\n'))
	return
}

// PrintThai print string text to Escpos p
// It does not wrap lines automatically
func PrintThai(p *escpos.Escpos, text string) {
	p.WriteRaw(scanUpper(text))
	p.WriteRaw(scanMiddle(text))
	p.WriteRaw(scanLower(text))
}
