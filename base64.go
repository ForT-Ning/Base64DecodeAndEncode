package main

import (
	"fmt"
	"log"
	//"unicode/utf8"
)

//A~Z : 0 ~ 25	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//a~z : 26 ~ 51 "abcdefghijklmnopqrstuvwxyz"
//0~9 : 52 ~ 61	"0123456789"
// + : 62		"+"
// / : 63		"/"

func RuneToChar(n rune) string {
	if 0 <= n && 25 >= n {
		return string('A' + n)
	} else if 26 <= n && 51 >= n {
		return string('a' + n - 26)
	} else if 52 <= n && 61 >= n {
		return string('0' + n - 52)
	} else if n == 62 {
		return "+"
	} else if n == 63 {
		return "/"
	}
	panic("some error int IntToChar")
}

func CharToRune(r rune) rune {
	if 'A' <= r && 'Z' >= r {
		return r - 'A'
	} else if 'a' <= r && 'z' >= r {
		return r - 'a' + 26
	} else if '0' <= r && '9' >= r {
		return r - '0' + 52
	} else if r == '+' {
		return 62
	} else if r == '/' {
		return 63
	}
	panic("some error int CharToInt")
}

//解码
func DecodeString(str string) string {
	var number int = 0x00
	var n16 int = 0x0ff
	var b []byte
	for ii, cc := range str {
		if cc == '=' {
			break
		}
		if ii%4 == 0 {
			number = int(CharToRune(cc))
			continue
		}
		number = number<<6 + int(CharToRune(cc))
		if number > 0x03fff {
			log.Println("'number' value error in EncodeString")
			return ""
		}

		new8 := number >> (6 - (ii%4)*2)

		number = (n16 >> ((ii%4 + 1) * 2)) & number

		b = append(b, byte(new8))
	}
	return string(b)
}

//加密
func EncodeString(b []byte) string {
	var n16 int = 0x0ff
	var nLast int = 0x00
	var strRet string = ""
	for ii, ss := range b {
		move := (ii%3)*2 + 2

		sum := nLast<<8 + int(ss)

		if sum > 0x0fff {
			log.Println("'sum' value error in EncodeString")
			return ""
		}

		temp := n16 >> (8 - move)
		nLast = temp & int(ss)

		number := sum >> move

		strRet += RuneToChar(rune(number))

		if (ii+1)%3 == 0 {
			strRet += RuneToChar(rune(nLast))
			nLast = 0x00
		}

		if ii == len(b)-1 {
			switch move {
			case 2:
				nLast = nLast << 4
				strRet += RuneToChar(rune(nLast)) + "=="
				break
			case 4:
				nLast = nLast << 2
				strRet += RuneToChar(rune(nLast)) + "="
				break
			case 6:
				break
			}
		}
	}
	//fmt.Println(strRet)
	return strRet
}

func main() {
	str := "我是谁"
	b := []byte(str)
	enStr := EncodeString(b)
	fmt.Println(enStr)
	deStr := DecodeString(enStr)
	fmt.Println(deStr)

}