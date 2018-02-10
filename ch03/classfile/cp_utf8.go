package classfile

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
import (
	"fmt"
	"unicode/utf16"
)

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func (self *ConstantUtf8Info) Str() string {
	return self.str
}

// 解析mutf8 -> utf16 -> utf32 -> string
func decodeMUTF8(bytes []byte) string {
	utflen := len(bytes)
	chars := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	charsCount := 0

	// 小于127的在这里解析
	for count < utflen {
		c = uint16(bytes[count])
		if c > 127 {
			break
		}
		count++
		chars[charsCount] = c
		charsCount++
	}

	// 大于127的在这里
	for count < utflen {
		c = uint16(bytes[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			count++
			chars[charsCount] = c
			charsCount++
		case 12, 13:
			// 这段magic得后面解释
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end!")
			}
			char2 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chars[charsCount] = c&0x1F<<6 | char2&0x3F
			charsCount++
		case 14:
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chars[charsCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charsCount++
		default:
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}

	// 最终的字符数组可能比utflen短
	chars = chars[0:charsCount]
	runes := utf16.Decode(chars)
	return string(runes)
}
