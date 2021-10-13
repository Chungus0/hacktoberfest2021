package main

import (
	"os"
)

var array []byte

func IsNumeric(s string) bool {
	alphabet := "1324657890-"
	a := []rune(alphabet)
	word := []rune(s)
	count := 0
	for i := 0; i < len(word); i++ {
		for k := 0; k < len(a); k++ {
			if word[i] == a[k] {
				count++
			}
		}
	}
	if count == len(word) {
		return true
	} else {
		return false
	}
}

func Atoi(s string) int {
	nmb := []byte(s)
	number := 0
	list := [10]int{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	num := 0
	count := 0
	countmin := 0
	check := 0
	if s != "" {
		if rune(nmb[0]) == rune('-') {
			nmb[0] = '0'
			check++
		}
		for i := range nmb {
			if rune(nmb[i]) == rune('-') && i != 0 {
				countmin++
			}
		}
		for i := range nmb {
			for k := 0; k < 10; k++ {
				if rune(nmb[i]) == rune(list[k]) {
					count++
				}
			}
		}
		for i := range nmb {
			if count != len(s) || countmin > 0 {
				num = 0
			} else {
				number = int(rune(nmb[i]) - 48)
				num = (num * 10) + number
			}
		}
		if check == 1 {
			num = num * (-1)
		}
	} else {
		num = 0
	}
	return num
}

func PrintNumber(n int) {
	d := n
	p := 1
	i := -1
	c := -1
	cn := 0
	if n < 0 {
		array = append(array, byte('-'))
		n = -n
		d = -d
	}
	if n != 0 {
		for d != 0 {
			d = d / 10
			i++
		}
		ii := i + 1
		for b := 0; b < ii; b++ {
			if i == 0 {
				p = 1
			} else {
				for k := 0; k < i; k++ {
					p = p * 10
				}
			}
			c = (n / p) - (cn * 10)
			cn = n / p
			p = 1
			i = i - 1
			array = append(array, 48+byte(c))
		}
	} else {
		array = append(array, 48)
	}
}

func main() {
	a := os.Args[1:]
	ops := []string{"+", "-", "/", "*", "%"}
	value := 0
	valid := -1
	text := ""
	if len(a) > 3 || len(a) < 3 {
		return
	} else {
		if IsNumeric(a[0]) && IsNumeric(a[2]) {
			for i := range ops {
				if a[1] == ops[i] {
					valid = 1
				}
			}
			if valid == 1 {
				num1 := Atoi(a[0])
				num2 := Atoi(a[2])
				if num1 >= 9223372036854775807 || num2 >= 9223372036854775807 || num1 <= -922337203685477580 || num1 <= -9223372036854775808 {
					return
				}
				if a[1] == ops[0] {
					value = num1 + num2
				}
				if a[1] == ops[1] {
					value = num1 - num2
				}
				if a[1] == ops[2] {
					if num2 == 0 {
						text = "No division by 0"
					} else {
						value = num1 / num2
					}
				}
				if a[1] == ops[3] {
					value = num1 * num2
				}
				if a[1] == ops[4] {
					if num2 == 0 {
						text = "No modulo by 0"
					} else {
						value = num1 % num2
					}
				}
				if text != "" {
					os.Stdout.Write([]byte(text))
					os.Stdout.Write([]byte("\n"))
				} else {
					PrintNumber(value)
					os.Stdout.Write(array)
					os.Stdout.Write([]byte("\n"))
				}
			}
		}
	}
}
